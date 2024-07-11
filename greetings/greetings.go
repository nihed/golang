package greetings

import (
  "fmt"
  "errors"
  "math/rand"
)
//Hello returns a greeting for the name person.

func Hello(name string) (string, error) {
  if name == "" {
    return "", errors.New("Empty name")
  }


  message := fmt.Sprintf(randomFormat(), name)
  return message,nil
}

func randomFormat() string {
  formats := []string{
   "Hi, %v. Welcome!",
   "Great to see you, %v!",
   "Hail, %v! Well met!",
  }
  return formats[rand.Intn(len(formats))]
}
