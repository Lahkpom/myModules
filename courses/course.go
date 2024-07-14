package courses

import "fmt"

type course struct {
	name    string
	price   float64
	isFree  bool
	UserIDs []uint
	// Classes map[uint]string
	classes []class
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
		name:   name,
		price:  price,
		isFree: isFree,
	}
}

// PrintClasses() Recibe un objeto Course e imprime el nombre y descripción de sus clases
func (c course) PrintClasses() {
	fmt.Println("Las clases de este curso son:")
	for _, class := range c.classes {
		fmt.Println(class.name)
		fmt.Println(class.description)
		fmt.Println("==================================")
	}
}

// ChangePrice() Recibe un valor y lo actualiza en el precio
func (c *course) ChangePrice(newPrice float64) {
	c.price = newPrice
}

// AddClass() Recibe name y description de una clase para crearla y añadirla a la lista de clases
func (c *course) AddClass(name string, description string) {
	newClass := NewClass(name, description)
	c.classes = append(c.classes, *newClass)
}
