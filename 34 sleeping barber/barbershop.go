package main

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity    int
	HaircutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientChan      chan string
	Open            bool
}

func (s *BarberShop) addBarber(barber string) {
	s.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for clients.", barber)
		for {
			// if there are no clients, barber goes to sleep
			if len(s.ClientChan) == 0 {
				color.Yellow("There is nothing to do, so %s takes a nap.", barber)
				isSleeping = true
			}
			client, shopOpen := <-s.ClientChan

			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes %s up.", client, barber)
					isSleeping = false
				}
				// then get a haircut
				s.haircut(barber, client)
			} else {
				// shop is closed, so send the barber home and close this goroutine
				s.sendBarberHome(barber)

				return
			}
		}
	}()
}

func (s *BarberShop) haircut(barber, client string) {
	color.Green("%s is cutting %s's hair.", barber, client)
	time.Sleep(s.HaircutDuration)
	color.Green("%s is finished cutting %s's hair.", barber, client)
}

func (s *BarberShop) sendBarberHome(barber string) {
	color.Cyan("%s is going home.", barber)
	s.BarbersDoneChan <- true
}

func (s *BarberShop) closeShopForDay() {
	color.Cyan("Closing shop for the day")

	close(s.ClientChan)
	s.Open = false

	for a := 1; a <= s.NumberOfBarbers; a++ {
		<-s.BarbersDoneChan
	}

	close(s.BarbersDoneChan)

	color.Green("---------------------------------------------------------------------")
	color.Green("The barbershop is now closed for the day, and everyone has gone home.")
}

func (s *BarberShop) addClient(client string) {
	// print out a message
	color.Green("*** %s arrives!", client)

	if s.Open {
		select {
		case s.ClientChan <- client:
			color.Yellow("%s takes a seat in the waiting room.", client)
		default:
			color.Red("The waiting room is full, so %s leaves.", client)
		}
	} else {
		color.Red("The shop is already closed, so %s leaves.", client)
	}
}
