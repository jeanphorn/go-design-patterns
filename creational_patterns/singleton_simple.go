/*
*
*
 */

package singleton

type SingleTon struct {
}

var instance *SingleTon

func GetInstance() *SingleTon {
	if Instance == nil {
		instance = &SingleTon{}
	}

	return instance
}
