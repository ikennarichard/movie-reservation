# 🎬 Movie Reservation Service

Welcome to the **Movie Reservation Service**! 🍿🎥

This is a backend system designed to help you manage movie reservations, browse the latest films and reserve seats for your favorite showtimes.

## 🌟 Features

- **User Authentication** 🔒  
  Sign up, log in, and manage your account. Admins can promote users and manage the system, while regular users can reserve seats!

- **Movie Management** 🎬  
  Admins can add, update, and remove movies. Movies are categorized by genre and each movie has showtimes!

- **Showtime Management** 🕒  
  Admins can schedule showtimes for movies, and users can reserve seats for these showtimes.

- **Reservation System** 🎟️  
  Reserve the best seats for your chosen movie! Keep track of your upcoming reservations and even cancel them if needed.

- **Admin Dashboard** 📊  
  Admins can view all reservations, check seat availability, and get revenue insights.

## 🚀 Getting Started

To get started with the Movie Reservation System locally, you'll need to set up your Go environment and a test database.

### 1. Clone the repo

```bash
git clone https://github.com/ikennarichard/movie-reservation.git
cd movie-reservation
```

### 2. Install Dependencies

Make sure you've got Go installed on your machine.

Install necessary Go dependencies:

```bash
go mod tidy
```

### 3. Run the Server

Now, let's start the server. Run:

```bash
go run .
```

The server should be running at [http://localhost:8080](http://localhost:8080). You can now make API requests!

### 4. API Endpoints

Here are the key routes available:

- **POST** `/signup` – Sign up a new user.
- **POST** `/login` – Log in with your credentials.
- **GET** `/movies` – Get a list of movies.
- **POST** `/reservations` – Reserve seats for a movie showtime.
- **GET** `/reservations` – View your current reservations.
- **DELETE** `/reservations/:id` – Cancel a reservation.

## 🛠️ Tech Stack

- **Go**
- **Gin**
- **GORM**
- **JWT**
- **PostgreSql**

---

## 💡 Future Improvements/ What I learned

- Add a **payment gateway integration** to handle payments for reservations.
- Enable **multi-theater support** for larger cinema chains.
- Updating a table with indexes takes more time than updating a table without (because the indexes also need an update). So, only create indexes on columns that will be frequently searched against.(W3Schools)

## 👨‍💻 Status

_In progress._

## 📝 Contributing

Contributions are welcome! If you have any ideas for improvements or bug fixes, feel free to submit a pull request. Here are some ways you can contribute:

1. **Report Bugs** – Found a bug? Let us know, and we'll work on fixing it.
2. **Add Features** – Got a cool feature in mind? Submit an issue or PR!
3. **Improve Documentation** – We always love better docs!

## 📞 Contact

For questions, suggestions, or if you just want to chat about movies 🎬, feel free to reach out to me via [email](mailto:oguefioforrichard@gmail.com) or raise an issue in the GitHub repository!
