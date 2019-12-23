package storage

import (
    pb "github.com/jonb377/website/storage-service/proto/storage"
    util "github.com/jonb377/website/router-service/router"
    "context"
    "github.com/jinzhu/gorm"
)

type StorageService struct {
    db *gorm.DB
}

func (s *StorageService) SaveBlob(ctx context.Context,  req *pb.Blob, resp *pb.Empty) error {
    session := util.GetSessionData(ctx)
    blob := Blob{
        User: session.Username,
        URI: session.MuxVars["uri"].(string),
        Data: req.Data,
        Modified: req.Date,
        DeletedAt: req.DeletedAt,
    }
    if err := s.db.Table("blobs").Set(
        "gorm:insert_option",
        "ON CONFLICT ON CONSTRAINT blobs_pkey DO UPDATE SET data = excluded.data, date = excluded.date, deleted = excluded.deleted",
        ).Create(&blob).Error; err != nil {
        return err
    }
    return nil
}

func (s *StorageService) Rename(ctx context.Context,  req *pb.Empty, resp *pb.Empty) error {
    session := util.GetSessionData(ctx)
    blob := Blob{
        User: session.Username,
        URI: session.MuxVars["olduri"].(string),
    }
    if err := s.db.Model(&blob).Update("uri", session.MuxVars["newuri"].(string)).Error; err != nil {
        return err
    }
    return nil
}

func (s *StorageService) DeleteBlob(ctx context.Context,  req *pb.Empty, resp *pb.Empty) error {
    session := util.GetSessionData(ctx)
    blob := Blob{
        User: session.Username,
        URI: session.MuxVars["uri"].(string),
    }
    if err := s.db.Delete(&blob).Error; err != nil {
        return err
    }
    return nil
}

func (s *StorageService) Sync(ctx context.Context,  req *pb.SyncRequest, resp *pb.SyncResponse) error {
    session := util.GetSessionData(ctx)
    var dbblobs []Blob
    if err := s.db.Table("blobs").Where("user = ? and date > ?", session.Username, req.LastUpdate).Find(&dbblobs).Error; err != nil {
        return err
    }
    resp.Blobs = make([]*pb.Blob, len(dbblobs))
    for i, p := range dbblobs {
        resp.Blobs[i] = &pb.Blob{
            Uri: p.URI,
            Date: p.Modified,
            Data: p.Data,
            DeletedAt: p.DeletedAt,
        }
    }
    return nil
}
