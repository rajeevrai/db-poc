package data_generator

import (
	"math/rand"
	"time"
)

const (
	base         = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	maxCeil      = 9999999999999
	firstJan2014 = 1388534400
)

func GenerateID() string {
	nanotime := time.Now().UTC().UnixNano()
	b62 := nanotimeToBase62(nanotime)

	rand.Seed(nanotime)

	random := int64(rand.Intn(maxCeil))
	base62Rand := base62(random)

	if len(base62Rand) > 4 {
		base62Rand = base62Rand[len(base62Rand)-4:]
	}

	return b62 + base62Rand
}

func base62(num int64) string {
	var res string
	index := base

	for {
		res = string(index[num%62]) + res
		num = int64(num / 62)

		if num == 0 {
			break
		}
	}

	return res
}

func nanotimeToBase62(nanotime int64) string {
	epochTs := int64(firstJan2014 * 1000000000)

	nanotime -= epochTs

	return base62(nanotime)
}
