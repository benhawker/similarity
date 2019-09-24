package main

import (
  	"fmt"
	"math"
)

type Film struct {
	Name string
	R1   int
	R2   int
}

type Rating struct {
	FilmName   string
	Score float64
}

type Critic struct {
	Name    string
	Ratings []Rating
}


func main() {
	ratings := []Rating{{"Lady in the Water", 2.5}, {"Snakes on a Plane", 3.5}, {"Just My Luck", 3.0}, {"Superman Returns", 3.5}, {"You, Me and Dupree", 2.5}, {"The Night Listener", 3.0}}
	ratings2 := []Rating{{"Lady in the Water", 3.0}, {"Snakes on a Plane", 3.5}, {"Just My Luck", 1.5}, {"Superman Returns", 5.0}, {"The Night Listener", 3.0}, {"You, Me and Dupree", 3.5}}

	critic1 := Critic{"Critic 1", ratings}
	critic2 := Critic{"Critic 2", ratings2}
	fmt.Println(pearsonCorrelation(critic1, critic2))
}

func pearsonCorrelation(critic1, critic2 Critic) float64 {

	matrix := buildMatrix(critic1, critic2)
	lengthOfMatrix := len(matrix)

	// Return 0 now if there are no films that both critics have rated
	if lengthOfMatrix == 0 {
		return 0
	}

	var sumOfCritic1Ratings, sumOfCritic2Ratings float64
	var sumOfSquaresForCritic1Ratings, sumOfSquaresForCritic2Ratings float64
	var sumOfProduct float64

	for _, film := range matrix {
		// Add up all the ratings for each critic
		sumOfCritic1Ratings += film.R1
		sumOfCritic2Ratings += film.R2

		// Sum up the squares for each critic
		critic1.Ratings[film.P1].Score
		
		sumOfSquaresForCritic1Ratings += math.Pow(float64(film.R1), 2)
		sumOfSquaresForCritic2Ratings += math.Pow(float64(film.R1), 2)

		// Sum up the products
		sumOfProduct += film.R1 * film.R2
	}

	// Calculate Pearson correlation score
	num := sumOfProduct - (sumOfCritic1Ratings * sumOfCritic2Ratings / lengthOfMatrix)
	preDen := (sumOfSquaresForCritic1Ratings - (math.Pow(sumOfCritic1Ratings, 2) / lengthOfMatrix)) * (sumOfSquaresForCritic2Ratings - (math.Pow(sumOfCritic2Ratings, 2) / lengthOfMatrix))
	den := math.Pow(preDen, 0.5)

	if den == 0.0 {
		return 0.0
	}

	return num/den
	// return 1 / (1 + math.Sqrt(sumOfSquares))

	// # Add up all the preferences
 	// sum1=sum([prefs[p1][it] for it in si])
	// sum2=sum([prefs[p2][it] for it in si])
	// # Sum up the squares
	// sum1Sq=sum([pow(prefs[p1][it],2) for it in si])
	// sum2Sq=sum([pow(prefs[p2][it],2) for it in si])
	// # Sum up the products
	// pSum=sum([prefs[p1][it]*prefs[p2][it] for it in si])
	// # Calculate Pearson score
	// num=pSum-(sum1*sum2/n)
	// den=sqrt((sum1Sq-pow(sum1,2)/n)*(sum2Sq-pow(sum2,2)/n))
	// if den==0: return 0
	// r=num/den
	// return r
}


func buildMatrix(critic1 Critic, critic2 Critic) []Film {
	var ratedFilms []Film

	for index, critic1Rating := range critic1.Ratings {
		for index2, critic2Rating := range critic2.Ratings {
			// If this film has been rated by both critics add it to the slice of films
			if critic1Rating.FilmName == critic2Rating.FilmName {
				ratedFilms = append(ratedFilms, Film{critic1Rating.FilmName, index, index2})
			}
		}
	}

	return ratedFilms
}
