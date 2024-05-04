package models

type User struct {
	Id              int    `db:"id"`
	Username        string `db:"username"`
	Phone           string `db:"phone"`
	Email           string `db:"email"`
	SubscriptionLvl int    `db:"subscription_lvl"`
	CurrentGymId    int    `db:"current_gym_id"`
}
