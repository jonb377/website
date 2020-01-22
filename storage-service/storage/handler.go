package storage

import (
    pb "github.com/jonb377/website/storage-service/proto/storage"
    util "github.com/jonb377/website/router-service/router"
    "context"
    "github.com/jinzhu/gorm"
    "time"
)

type StorageService struct {
    db *gorm.DB
}

func (s *StorageService) SaveBlob(ctx context.Context,  req *pb.Blob, resp *pb.Empty) error {
    session := util.GetSessionData(ctx)
    blob := Blob{
        User: session.Username,
        URI: req.Uri,
        Data: req.Data,
        Modified: req.Date,
    }
    if req.DeletedAt != 0 {
        blob.DeletedAt = time.Unix(req.DeletedAt, 0)
    }
    if err := s.db.Table("blobs").Set(
        "gorm:insert_option",
        "ON CONFLICT ON CONSTRAINT blobs_pkey DO UPDATE SET data = excluded.data, modified = excluded.modified, deleted_at = excluded.deleted_at",
        ).Create(&blob).Error; err != nil {
        return err
    }
    return nil
}

func (s *StorageService) Rename(ctx context.Context,  req *pb.RenameRequest, resp *pb.Empty) error {
    session := util.GetSessionData(ctx)
    blob := Blob{
        User: session.Username,
        URI: req.OldUri,
    }
    if err := s.db.Model(&blob).Update("uri", req.NewUri).Error; err != nil {
        return err
    }
    return nil
}

func (s *StorageService) DeleteBlob(ctx context.Context,  req *pb.Blob, resp *pb.Empty) error {
    session := util.GetSessionData(ctx)
    blob := Blob{
        User: session.Username,
        URI: req.Uri,
    }
    if err := s.db.Delete(&blob).Error; err != nil {
        return err
    }
    return nil
}

func (s *StorageService) Sync(ctx context.Context,  req *pb.SyncRequest, resp *pb.SyncResponse) error {
    session := util.GetSessionData(ctx)
    var dbblobs []Blob
    if err := s.db.Table("blobs").Where("user = ? and modified > ?", session.Username, req.LastUpdate).Find(&dbblobs).Error; err != nil {
        return err
    }
    resp.Blobs = make([]*pb.Blob, len(dbblobs))
    for i, p := range dbblobs {
        resp.Blobs[i] = &pb.Blob{
            Uri: p.URI,
            Date: p.Modified,
            Data: p.Data,
        }
        if p.DeletedAt.IsZero() {
            resp.Blobs[i].DeletedAt = p.DeletedAt.Unix()
        }
    }
    return nil
}
