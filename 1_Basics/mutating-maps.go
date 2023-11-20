// Задача: Создание Простой Телефонной Книги на Go

// 1. Определите Структуру для Контактов:
// Создайте структуру с именем Contact, которая включает поля Name (строка), Email (строка) и Phone (строка).
// 2. Создайте Слайс Контактов: Создайте слайс для хранения нескольких контактов.
// 3. Добавьте Контакты в Слайс: Напишите функцию для добавления новых контактов в ваш слайс.
// Эта функция должна принимать имя, электронную почту и телефон в качестве параметров,
// а затем добавлять новый Contact в ваш слайс.
// 4. Поиск Контакта: Реализуйте функцию, которая принимает имя в качестве аргумента и ищет это имя в вашем слайсе контактов.
// Если контакт найден, функция должна выводить детали контакта; если нет, то выводить "Контакт не найден".
// 5. Дополнительно - Реализуйте Использование Map: Вместо слайса используйте map для хранения ваших контактов.
// Ключом должно быть имя контакта, а значением - структура Contact.
// Измените ваши функции добавления и поиска так, чтобы они работали с этой map.

package main

import "fmt"

type Contact struct {
	Name  string
	Email string
	Phone string
}

func main() {
	contacts := map[string]Contact{}
	fmt.Println(contacts)

	AddContact("Bulat", "bulat@mail.ru", "911", &contacts)
	fmt.Println(contacts)

	fmt.Println(FindContact("Bulat", contacts))
	fmt.Println(FindContact("Denis", contacts))

}

func AddContact(name, email, phone string, contactMap *map[string]Contact) {
	(*contactMap)[name] = Contact{name, email, phone}
}

func FindContact(name string, contactMap map[string]Contact) string {
	contact, ok := contactMap[name]
	if ok {
		return contact.Name + " " + contact.Email + " " + contact.Phone
	}

	return "Контакт не найден"
}
