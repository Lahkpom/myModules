package courses

import "fmt"

type course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	// Classes map[uint]string
	Classes []Class
}

// El constructor devuelve un puntero para que pueda utilizarse en la versión original para efectuar cambios, sino
// es solo una copia

// NewCourse() Constructor, devuelve un puntero a un curso
func NewCourse(name string, price float64, isFree bool) (a *course) {
	// Para controlar lo valores por defecto lo controlamos antes del return
	if price <= 0 {
		price = 100
	}

	// Para devolver un puntero tengo que ponerle el &
	// Esto crea el objeto y se lo cevuelve para la variable que lo recibe en el otro lado
	// que lo está solicitando
	return &course{
		Name:    name,
		Price:   price,
		IsFree:  isFree,
		UserIDs: make([]uint, 0),
	}
}

// PrintClasses() Recibe un objeto Course e imprime el nombre y descripción de sus clases
func (c course) PrintClasses() {
	fmt.Println("Las clases de este curso son:")
	for _, class := range c.Classes {
		fmt.Println(class.Name)
		fmt.Println(class.Description)
		fmt.Println("==================================")
	}
}

// ChangePrice() Recibe un valor y lo actualiza en el precio
func (c *course) ChangePrice(newPrice float64) {
	c.Price = newPrice
}
