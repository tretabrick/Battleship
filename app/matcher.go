package matcher

import (
	"errors"
	"math/rand"
)

type matcher struct {
	IPtoCode map[string]int // associate IP with game invite code
	CodetoIP map[int]string
}

func (m matcher) New() matcher {
	m = matcher{
		IPtoCode: make(map[string]int),
		CodetoIP: make(map[int]string),
	}
	return m
}
func (m matcher) getCode(ip string) int {
	code, ok := m.IPtoCode[ip]
	if !ok {
		for {
			newcode := rand.Int()
			_, ok := m.CodetoIP[newcode]

			if !ok {
				continue
			}
			code = newcode
			m.IPtoCode[ip] = code
			m.CodetoIP[code] = ip
			break
		}
	}
	return code
}

func (c matcher) iniciateGame(ownIP string, code int) (string, string, error) {

	var err error = nil
	ip1 := ownIP
	ip2, ok := c.CodetoIP[code]
	if !ok {
		err = errors.New("Code doesnt exist")
	}
	if ip1 == ip2 {
		err = errors.New("The code corresponds to the same ip")
	}

	if err == nil {

		ownCode := c.IPtoCode[ownIP]
		delete(c.CodetoIP, ownCode)
		delete(c.IPtoCode, ownIP)
		return ip1, ip2, err

	} else {
		return ip1, ip2, err
	}

}
