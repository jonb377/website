package auth

import (
    "math/big"
    "crypto/sha256"
    "crypto/rand"
)

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func equal(a, b []byte) bool {
    if len(a) != len(b) {
        return false
    }
    for i, ai := range a {
        if b[i] != ai {
            return false
        }
    }
    return true
}

var N *big.Int
var g *big.Int
var hNxorg *big.Int
var k *big.Int
var NBytes []byte
var gBytes []byte

func init() {
    N, _ = big.NewInt(0).SetString("3195175819128742910148899061986555150121416420444122774875756046" +
                "30571647693625360441888077057362255972535071846348545922743045172817543421299842516366630380803616622971" +
                "71169726698521325900924027392844935451872772319366851484863111798246795520996654353488237194689495898860" +
                "37102934582600605020436652165266026373297760632067386811657666094511049691456531848155483223265114552859" +
                "37694157993515058640163701505911137152029179449178571830500045815764530877766080261063507153499708634615" +
                "63227435927003652822470523123962910966326470801300421844308747559664405151192019435188478751182102168045" +
                "041793771480326644387019638208599", 10)

    g, _ = big.NewInt(0).SetString("1348291216488854302997516880256913169207251411704710281765062907" +
                "99803287345361635469013674530171333302961590309733624526554294085754083242340156278307798812793066128743" +
                "36672278135267192067947445838661167209263820873612240071618265961340761546498787123671356182068349958400" +
                "3629920530928740801642934388698947533", 10)

    NBytes := N.Bytes()
    gBytes := g.Bytes()
    paddedg := pad(len(NBytes), gBytes)
    hash := sha256.Sum256(append(NBytes, paddedg...))
    k = big.NewInt(0).SetBytes(hash[:])
    NHash := sha256.Sum256(NBytes)
    gHash := sha256.Sum256(paddedg)
    for i, gi := range gHash {
        NHash[i] ^= gi
    }
    hNxorg = big.NewInt(0).SetBytes(NHash[:])
}

func generate_b() []byte {
    minBits := uint(min(256, N.BitLen() / 2))
    min := big.NewInt(0).Lsh(big.NewInt(1), minBits)
    max := big.NewInt(0).Sub(N, big.NewInt(1))
    diff := big.NewInt(0).Sub(max, min)
    n, err := rand.Int(rand.Reader, diff)
    if err != nil {
        panic("Failed to generate b")
    }
    // Move back into the range [min, max)
    return n.Add(n, min).Bytes()
}

func pad(targetLen int, b []byte) []byte {
    if len(b) < targetLen {
        return append(make([]byte, targetLen - len(b)), b...)
    }
    return b
}

type SRPServer struct {
    A *big.Int          // Client ephemeral key
    b *big.Int          // Ephemeral key generator
    B *big.Int          // Server ephemeral key
    s []byte            // Salt
    v *big.Int          // Verifier
    u *big.Int          // Hash of A and B
    S *big.Int          // Shared secret
    K []byte            // The key to be used
    M []byte            // Client verifier
    HAMK []byte         // Hash of A, M, and K
    username string     // Username
    authenticated bool  // Has the authentication succeeded?
}

func NewSRPServerWithB(username string, s []byte, v []byte, ABytes []byte, bBytes []byte) *SRPServer {
    server := new(SRPServer)
    server.s = s
    server.A = big.NewInt(0).SetBytes(ABytes)
    if big.NewInt(0).Mod(server.A, N).Cmp(big.NewInt(0)) == 0 {
        // Safety check failed
        return nil
    }
    server.v = big.NewInt(0).SetBytes(v)
    server.b = big.NewInt(0).SetBytes(bBytes);
    server.B = big.NewInt(0).Mod(big.NewInt(0).Add(big.NewInt(0).Mul(k, server.v), big.NewInt(0).Exp(g, server.b, N)), N)
    server.username = username
    return server
}

func NewSRPServer(username string, s []byte, v []byte, ABytes []byte) *SRPServer {
    return NewSRPServerWithB(username, s, v, ABytes, generate_b())
}

// Verifies the client's M and returns the HAMK. If verification fails, returns nil
func (server *SRPServer) VerifyM(M []byte) []byte {
    server.u = server.calculate_u()
    if server.u.Cmp(big.NewInt(0)) == 0 {
        // Safety check failed
        return nil
    }
    server.S = big.NewInt(0).Exp(big.NewInt(0).Mul(server.A, big.NewInt(0).Exp(server.v, server.u, N)), server.b, N)
    hash := sha256.Sum256(server.S.Bytes())
    server.K = hash[:]
    server.M = server.calculateM()
    if !equal(server.M, M) {
        return nil
    }
    server.authenticated = true
    server.HAMK = server.calculateHAMK()
    return server.HAMK
}

func (s *SRPServer) getKey() []byte {
    return s.K
}

func (s *SRPServer) calculate_u() *big.Int {
    paddedA := pad(len(NBytes), s.A.Bytes())
    paddedB := pad(len(NBytes), s.B.Bytes())
    hash := sha256.Sum256(append(paddedA, paddedB...))
    return big.NewInt(0).SetBytes(hash[:])
}

func (s *SRPServer) calculateM() []byte {
    I := sha256.Sum256([]byte(s.username))
    h := sha256.New()
    h.Write(hNxorg.Bytes())
    h.Write(I[:])
    h.Write(s.s)
    h.Write(s.A.Bytes())
    h.Write(s.B.Bytes())
    h.Write(s.K)
    return h.Sum(nil)
}

func (s *SRPServer) calculateHAMK() []byte {
    h := sha256.New()
    h.Write(s.A.Bytes())
    h.Write(s.M)
    h.Write(s.K)
    return h.Sum(nil)
}
