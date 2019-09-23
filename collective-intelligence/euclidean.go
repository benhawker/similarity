package collective

// # A dictionary of movie critics and their ratings of a small
// # set of movies
// critics={'Lisa Rose': {'Lady in the Water': 2.5, 'Snakes on a Plane': 3.5,
//  'Just My Luck': 3.0, 'Superman Returns': 3.5, 'You, Me and Dupree': 2.5,
//  'The Night Listener': 3.0},
// 'Gene Seymour': {'Lady in the Water': 3.0, 'Snakes on a Plane': 3.5,
//  'Just My Luck': 1.5, 'Superman Returns': 5.0, 'The Night Listener': 3.0,
//  'You, Me and Dupree': 3.5},
// 'Michael Phillips': {'Lady in the Water': 2.5, 'Snakes on a Plane': 3.0,
//  'Superman Returns': 3.5, 'The Night Listener': 4.0},
// 'Claudia Puig': {'Snakes on a Plane': 3.5, 'Just My Luck': 3.0,
//  'The Night Listener': 4.5, 'Superman Returns': 4.0,
//  'You, Me and Dupree': 2.5},
// 'Mick LaSalle': {'Lady in the Water': 3.0, 'Snakes on a Plane': 4.0,
//  'Just My Luck': 2.0, 'Superman Returns': 3.0, 'The Night Listener': 3.0,
//  'You, Me and Dupree': 2.0},
// 'Jack Matthews': {'Lady in the Water': 3.0, 'Snakes on a Plane': 4.0,
//  'The Night Listener': 3.0, 'Superman Returns': 5.0, 'You, Me and Dupree': 3.5},
// 'Toby': {'Snakes on a Plane':4.5,'You, Me and Dupree':1.0,'Superman Returns':4.0}}

import (
  	"fmt"
	"math"
)

type Film struct {
	Film string
	P1   int
	P2   int
}

type Rating struct {
	FilmName   string
	Score float64
}

type Critic struct {
	Name    string
	Ratings []Rating
}

func collaborative() {
	ratings := []Rating{{"Lady in the Water", 2.5}, {"Snakes on a Plane", 3.5}, {"Just My Luck", 3.0}, {"Superman Returns", 3.5}, {"You, Me and Dupree", 2.5}, {"The Night Listener", 3.0}}	// ratings2 := []Rating{{"Lady in the Water", 2.5}, {"Snakes on a Plane", 3.5}, {"Just My Luck", 3.0}, {"Superman Returns", 3.5}, {"You, Me and Dupree", 2.5}, {"The Night Listener", 3.0}}
	ratings2 := []Rating{{"Lady in the Water", 3.0}, {"Snakes on a Plane", 3.5}, {"Just My Luck", 1.5}, {"Superman Returns", 5.0}, {"The Night Listener", 3.0}, {"You, Me and Dupree", 3.5}}

	critic1 := Critic{"Lisa Rose", ratings}
	critic2 := Critic{"Gene Seymour", ratings2}
	fmt.Println(euclideanDistance(critic1, critic2))
}

func euclideanDistance(critic1, critic2 Critic) float64 {

	matrix := buildMatrix(critic1, critic2)

	// Return 0 now if there are no films that both critics have rated
	if len(matrix) == 0 {
		return 0
	}

	var sumOfSquares float64

	for _, film := range matrix {
		critic1MinusCritic2Score := critic1.Ratings[film.P1].Score - critic2.Ratings[film.P2].Score
		sumOfSquares += math.Pow(critic1MinusCritic2Score, 2)
	}

	return 1 / (1 + math.Sqrt(sumOfSquares))
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
