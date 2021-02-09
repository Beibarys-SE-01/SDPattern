package main

import (
	"fmt"
	"log"
	s "strings"
)

func qwe(){
	//in this case error will not appear
	movieFacade := newMovieFacade("Never give up", 10000000)
	err := movieFacade.changeReleaseDateForMovie("Never give up", "2020/10/04")
	if err != nil {
		log.Fatal("Error: ", err.Error())
	}
	fmt.Println("____________________________________________")
	//in this case error will appear, because name is incorrect
	movieFacade = newMovieFacade("Never give up", 10000000)
	err = movieFacade.changeReleaseDateForMovie("Never give ", "2020/10/04")
	if err != nil {
		log.Fatal("Error: ", err.Error())
	}
	fmt.Println("____________________________________________")
	//in this case error will appear, because date format is incorrect
	movieFacade = newMovieFacade("Never give up", 10000000)
	err = movieFacade.changeReleaseDateForMovie("Never give up", "2020-10-04")
	if err != nil {
		log.Fatal("Error: ", err.Error())
	}
}

type movieFacade struct {
	movie *movie
	money *money
	releaseDate * releaseDate
}

func newMovieFacade(movieName string, cash int) *movieFacade {
	fmt.Println("The plot for the new film is being discussed")
	movieFacade := &movieFacade{
		movie: newMovie(movieName),
		money: newMoney(cash),
		releaseDate: newReleaseDate(),
	}
	fmt.Println("Congratulations, the plot of your film has been agreed")
	return movieFacade
}

func (mf *movieFacade) changeReleaseDateForMovie(movieName string, time string) error {
	err := mf.movie.checkMovieName(movieName)
	if err != nil {
		return err
	}
	err = mf.releaseDate.setReleaseDate(time)
	if err != nil {
		return err
	}
	return nil
}


type movie struct {
	movieName string
}

func newMovie(movieName string) *movie {
	return &movie{
		movieName: movieName,
	}
}

func (a *movie) checkMovieName(movieName string) error {
	if a.movieName != movieName {
		return fmt.Errorf("Movie name is incorrect")
	}
	fmt.Println("Movie has been found")
	return nil
}

type money struct {
	cash int
}

func newMoney(cash int) *money {
	return &money{
		cash: cash,
	}
}

type releaseDate struct {
	releaseDate string
}

func newReleaseDate() *releaseDate {
	return &releaseDate{
		releaseDate: "",
	}
}

func (w *releaseDate) setReleaseDate(time string) error {
	err := w.checkReleaseDate(time)
	if err != nil {
		w.releaseDate = time
		fmt.Println("Date of release changed!")
	}
	return err
}

func (w *releaseDate) checkReleaseDate(time string) error {
	if !s.Contains(time, "/") {
		return fmt.Errorf("Your date format is incorrect, use / ")
	}
	fmt.Println("Your date format is fine")
	return nil
}