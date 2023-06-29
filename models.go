package main

import (
	"time"

	"gorm.io/gorm"
)

type Miner struct {
	gorm.Model
	Address                string `gorm:"size:255;uniqueIndex"`
	LastNotification       time.Time
	LastNotificationWeekly time.Time `gorm:"default:'2023-06-17 23:00:00.797487649+00:00'"`
	TelegramId             int64
	MiningHeight           int64
	MiningTime             time.Time
	ReferralID             uint `gorm:"index"`
	Confirmed              bool `gorm:"default:false"`
	Balance                uint64
	MinedTelegram          uint64
	MinedMobile            uint64
	LastPing               time.Time
	PingCount              int64
	// IpAddresses            []*IpAddress `gorm:"many2many:miner_ip_addresses;"`
	UpdatedApp          bool `gorm:"default:false"`
	LastInvite          time.Time
	BatteryNotification bool `gorm:"default:false"`
}
