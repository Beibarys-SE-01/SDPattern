package week3

import "fmt"

type Movie struct {
	//Builder 1 - MoviePlanBuilder
	nameOfTheMovie, typeOfMovie, producer, genre, mainCharacter string
	//by type of the movie, I mean that, it can be a cartoon or serial

	//Builder 2 - MovieLocationBuilder
	country,city string //filming locations

	//Builder 3 - MovieCalculationBuilder
	appYearOfRelease,appFee string
	//approximately year of movie's release;approximately amount of money, which can be earned after release
}

//Main Builder
type MovieBuilder struct {
	movie *Movie
}

//Constructor of the main builder - MovieBuilder
func NewMovieBuilder() *MovieBuilder {
	return &MovieBuilder{&Movie{}}
}

func (mb *MovieBuilder) Build() *Movie {
	return mb.movie
}

//First function of the main builder - MovieBuilder
func (mb *MovieBuilder) Plan() *MoviePlanBuilder {
	return &MoviePlanBuilder{*mb}
}

//Builder 1 - MoviePlanBuilder
type MoviePlanBuilder struct {
	MovieBuilder
}

func (pb *MoviePlanBuilder) SetNameOfTheMovie(name string) *MoviePlanBuilder{
	pb.movie.nameOfTheMovie = name
	return pb
}

func (pb *MoviePlanBuilder) SetTypeOfMovie(name string) *MoviePlanBuilder{
	pb.movie.typeOfMovie = name
	return pb
}

func (pb *MoviePlanBuilder) SetProducer(nameOfProducer string) *MoviePlanBuilder {
	pb.movie.producer = nameOfProducer
	return pb
}

func (pb *MoviePlanBuilder) SetGenre(genre string) *MoviePlanBuilder {
	pb.movie.genre = genre
	return pb
}

func (pb *MoviePlanBuilder) SetMainCharacter(nameOfActor string) *MoviePlanBuilder {
	pb.movie.mainCharacter = nameOfActor
	return pb
}


//Second function of the main builder - MovieBuilder
func (mb *MovieBuilder) SetLocation() *MovieLocationBuilder {
	return &MovieLocationBuilder{*mb}
}

//Builder 2 - MovieLocationBuilder
type MovieLocationBuilder struct {
	MovieBuilder
}

func (lb *MovieLocationBuilder) SetCountry(countryName string) *MovieLocationBuilder {
	lb.movie.country = countryName
	return lb
}

func (lb *MovieLocationBuilder) SetCity(cityName string) *MovieLocationBuilder {
	lb.movie.city = cityName
	return lb
}


func (mb *MovieBuilder) Calculation() *MovieCalculationBuilder {
	return &MovieCalculationBuilder{*mb}
}

type MovieCalculationBuilder struct {
	MovieBuilder
}

func (cb *MovieCalculationBuilder) SetYearOfRelease(year string) *MovieCalculationBuilder {
	cb.movie.appYearOfRelease = year
	return cb
}

func (cb *MovieCalculationBuilder) SetFee(fee string) *MovieCalculationBuilder {
	cb.movie.appFee = fee
	return cb
}


func builder() {
	movieBuilder := NewMovieBuilder()
	movieBuilder.Plan().
		SetTypeOfMovie("Film").
		SetGenre("Comedy").
		SetMainCharacter("Jim Carrey").
		SetProducer("Jim Carrey").
		SetLocation().
		SetCountry("Kazakhstan").
		SetCity("Nur-Sultan").
		Calculation().
		SetYearOfRelease("Never").
		SetFee("0$")
	movie := movieBuilder.Build()
	fmt.Println(*movie)
}

