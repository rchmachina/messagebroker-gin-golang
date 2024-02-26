package helper

import ("strings"

		"math/rand"
		"time"
)


func RandomWord()string{
	Word:= "Lorem ipsum dolor sit amet consectetur adipiscing elit Donec lacinia libero at sodales imperdiet nulla diam tincidunt mi in feugiat quam erat ut lorem Quisque sollicitudin elit vitae gravida viverra ante turpis lobortis ligula non ultrices ipsum est eu libero. Phasellus porttitor metus scelerisque eleifend justo et vulputate elit. Sed sollicitudin enim id nisl sollicitudin eleifend. Mauris et tincidunt lectus eget maximus lacus. Quisque commodo dui ac ultrices fringilla est tellus luctus enim ac pharetra libero ex vitae enim. In varius blandit tortor fermentum vehicula lorem condimentum sit amet. Duis diam erat laoreet pellentesque porta sed aliquam id enim. Etiam sit amet enim non quam volutpat suscipit. Sed tincidunt aliquet facilisis. Sed euismod mattis orci ac pellentesque"
	randomStr :=  strings.Split(Word," ")

	rand.Seed(time.Now().UnixNano())

    // Generate a random number in the range [0, 9999]
    randomNumber := rand.Intn(10000)
	return randomStr[randomNumber]
}