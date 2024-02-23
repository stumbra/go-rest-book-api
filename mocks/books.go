package mocks

import "go-rest-book-api/models"

var MockedBooks = []models.Book{

	{
		ID:          "1",
		Title:       "The Great Gatsby",
		Description: "The story of the mysteriously wealthy Jay Gatsby and his love for the beautiful Daisy Buchanan.",
		Author:      "F. Scott Fitzgerald",
		Published:   1925,
		InStore:     false,
	},
	{
		ID:          "2",
		Title:       "To Kill a Mockingbird",
		Description: "The story of a young girl in a Southern town and her father, a lawyer who defends a black man accused of raping a white woman.",
		Author:      "Harper Lee",
		Published:   1960,
		InStore:     true,
	},
	{
		ID:          "3",
		Title:       "1984",
		Description: "A dystopian novel by George Orwell, and one of the best-known works of dystopian fiction.",
		Author:      "George Orwell",
		Published:   1949,
		InStore:     false,
	},
	{
		ID:          "4",
		Title:       "Moby-Dick",
		Description: "The story of the voyage of the whaling ship Pequod, commanded by Captain Ahab, who leads his crew on a hunt for the great whale Moby Dick.",
		Author:      "Herman Melville",
		Published:   1851,
		InStore:     true,
	},
	{
		ID:          "5",
		Title:       "Pride and Prejudice",
		Description: "The story follows the main character, Elizabeth Bennet, as she deals with issues of manners, upbringing, morality, education, and marriage in the society of the landed gentry of early 19th-century England.",
		Author:      "Jane Austen",
		Published:   1813,
		InStore:     true,
	},
	{
		ID:          "6",
		Title:       "The Catcher in the Rye",
		Description: "The story of Holden Caulfield's experiences in New York City in the days following his expulsion from Pencey Prep, a fictional college preparatory school.",
		Author:      "J.D. Salinger",
		Published:   1951,
		InStore:     true,
	},
	{
		ID:          "7",
		Title:       "Harry Potter and the Philosopher's Stone",
		Description: "The story of a young wizard, Harry Potter, and his friends Hermione Granger and Ron Weasley, who are students at Hogwarts School of Witchcraft and Wizardry.",
		Author:      "J.K. Rowling",
		Published:   1997,
		InStore:     true,
	},
	{
		ID:          "8",
		Title:       "The Lord of the Rings",
		Description: "The story begins in the Shire, where the hobbit Frodo Baggins inherits the Ring from Bilbo Baggins, his cousin and guardian.",
		Author:      "J.R.R. Tolkien",
		Published:   1954,
		InStore:     true,
	},
	{
		ID:          "9",
		Title:       "The Hobbit",
		Description: "The story of Bilbo Baggins, a hobbit who lives in Hobbiton, who is visited by the wizard Gandalf and 13 dwarves.",
		Author:      "J.R.R. Tolkien",
		Published:   1937,
		InStore:     true,
	},
	{
		ID:          "10",
		Title:       "The Adventures of Huckleberry Finn",
		Description: "The story of Huck Finn's adventures with his friend Tom Sawyer, and their journey down the Mississippi River on a raft.",
		Author:      "Mark Twain",
		Published:   1884,
		InStore:     true,
	},
}
