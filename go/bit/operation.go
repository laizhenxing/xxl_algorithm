package bit

type Operation interface {
	Addition(int, int) int
	Subtraction(int, int) int
	Multiplication(int, int) int
	Division(int, int) int
}
