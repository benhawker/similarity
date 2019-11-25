package main

import (
  	"fmt"
	"math"
)

type Film struct {
	Name string
	R1   float64
	R2   float64
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
	ratings := []Rating{{"Lady in the Water", 2.5}, {"Snakes on a Plane", 3.5}, {"Just My Luck", 3.0}, {"Superman Returns", 3.5}, {"The Night Listener", 3.0}, {"You, Me and Dupree", 2.5}}
	ratings2 := []Rating{{"Lady in the Water", 3.0}, {"Snakes on a Plane", 3.5}, {"Just My Luck", 1.5}, {"Superman Returns", 5.0}, {"The Night Listener", 3.0}, {"You, Me and Dupree", 3.5}}

	critic1 := Critic{"Critic 1", ratings}
	critic2 := Critic{"Critic 2", ratings2}

	fmt.Printf("Pearson coefficient between critic 1 and critic 2: %f \n", pearsonCorrelation(critic1, critic2))
}

func pearsonCorrelation(critic1, critic2 Critic) float64 {
	matrix := buildMatrix(critic1, critic2)
	lengthOfMatrix := float64(len(matrix))

	// Return 0 now if there are no films that both critics have rated
	if lengthOfMatrix == 0.0 {
		return 0.0
	}

	var sumOfCritic1Ratings, sumOfCritic2Ratings float64
	var sumOfSquaresForCritic1Ratings, sumOfSquaresForCritic2Ratings float64
	var sumOfProduct float64

	for _, film := range matrix {
		// Add up all the ratings for each critic
		sumOfCritic1Ratings += film.R1
		sumOfCritic2Ratings += film.R2

		// Sum up the squares for each critic
		sumOfSquaresForCritic1Ratings += math.Pow(film.R1, 2)
		sumOfSquaresForCritic2Ratings += math.Pow(film.R2, 2)

		// Sum up the products
		sumOfProduct += film.R1 * film.R2
	}

	// Calculate Pearson correlation score
	num := sumOfProduct - (sumOfCritic1Ratings * sumOfCritic2Ratings / lengthOfMatrix) 
	preDen := (sumOfSquaresForCritic1Ratings - (math.Pow(sumOfCritic1Ratings, 2) / lengthOfMatrix)) * (sumOfSquaresForCritic2Ratings - (math.Pow(sumOfCritic2Ratings, 2) / lengthOfMatrix))
	den := math.Sqrt(preDen)

	if den == 0.0 {
		return 0.0
	}

	return num / den
}


func findMostSimilarReviewerTo(critic Critic) {
	for critic, _ : range allRatings {
		critic
	}

	


}

 // # Returns the closest entities from a given entity. The distance
 //    # between entities is based on the Pearson correlation coefficient
 //    #
 //    # @param [Hash] Hash containing entity-item scores
 //    # @param [String] Entity
 //    # @param [Hash] Options (limit)
 //    #
 //    # @return [Array] Top matches
 //    def closest_entities(scores, entity, opts={})
 //      sort_desc(scores, opts) do |h, k, v|
 //        entity == k ? h : h.merge(k => coefficient(scores, entity, k))
 //      end 
 //    end

 //    # Returns the best recommended items for a given entity
 //    #
 //    # @param [Hash] Hash containing entity-item scores
 //    # @param [String] Entity
 //    # @param [Hash] Options (limit)
 //    #
 //    # @return [Array] Top matches [item, score]
 //    def recommendations(scores, entity, opts={})
 //      totals = {}
 //      similarity_sums = {}

 //      totals.default = 0
 //      similarity_sums.default = 0

 //      fail EntityNotFound unless scores[entity]

 //      scores.each do |other_entity|
 //        next if other_entity.first == entity

 //        similarity = coefficient(scores, entity, other_entity.first)

 //        next if similarity <= 0

 //        scores[other_entity.first].each do |item, score|
 //          if !scores[entity].keys.include?(item) || scores[entity][item] == 0
 //            totals[item] += score * similarity
 //            similarity_sums[item] += similarity
 //          end
 //        end
 //      end

 //      sort_desc(totals, opts) {|h, k, v| h.merge(k => v/similarity_sums[k]) }
 //    end

 //    private

 //    # Returns a hash containing the shared items between two different entities
 //    #
 //    # @param [Hash] Hash containing entity-item scores
 //    # @param [String] Entity
 //    # @param [String] Entity
 //    #
 //    # @return [Hash] Common items
 //    def shared_items(scores, entity1, entity2)
 //      Hash[*(scores[entity1].keys & scores[entity2].keys).flat_map{|k| [k, 1]}]
 //    end

 //    def sort_desc(results, opts={})
 //      limit = opts[:limit] || 3

 //      results.reduce({}) do |h, (k, v)|
 //        yield(h, k, v)
 //      end.sort_by{|k, v| v}.reverse[0..(opts[:limit] || 3)-1]
 //    end
 //  end

// # Returns the best matches for person from the prefs dictionary.
// # Number of results and similarity function are optional params.
// def topMatches(prefs,person,n=5,similarity=sim_pearson):
//  scores=[(similarity(prefs,person,other),other)
//  for other in prefs if other!=person]
//  # Sort the list so the highest scores appear at the top
//  scores.sort( )
//  scores.reverse( )
//  return scores[0:n]

func buildMatrix(critic1 Critic, critic2 Critic) []Film {
	var ratedFilms []Film

	for _, critic1Rating := range critic1.Ratings {
		for _, critic2Rating := range critic2.Ratings {
			// If this film has been rated by both critics add it to the slice of films with the coresponding scores
			if critic1Rating.FilmName == critic2Rating.FilmName {
				ratedFilms = append(ratedFilms, Film{critic1Rating.FilmName, critic1Rating.Score, critic2Rating.Score})
			}
		}
	}

	return ratedFilms
}
