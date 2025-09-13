package models

import "time"

type User struct {
	UserID       int
	FirstName    string
	LastName     string
	Email        string
	PasswordHash string
}

type Trip struct {
	TripID        int
	DestinationID int
	Title         string
	Description   string
	Price         float64
	StartDate     time.Time
	EndDate       time.Time
}

type Booking struct {
	BookingID int
	UserID int
	TripID int
	BookingDate time.Time
	Status string
}

type Payment struct {
	PaymentID int
	BookingID int
	Amount float64
	PaymentDate time.Time
	Status string
}

type Notification struct {
	NotificationID int
	UserID int
	Message string
	IsRead bool
}

type Review struct {
	ReviewID int
	UserID int
	TripID int
	Rating int
	Comment string
}

type Destination struct {
	DestinationID int
	Country string
	City string
	Description string
}

type UserContact struct {
	ContactID int
	UserID int
	PhoneNumber string
	Address string
}

