package main

type IArea interface {
	getArea() (error, float64)
}