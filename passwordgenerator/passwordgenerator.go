package passwordgenerator

import (
	"crypto/rand"
	"crypto/sha1"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strings"
)

func GeneratePassword(minLen int, maxLen int, lowLetter bool, bigLetter bool, number bool, specialCharacter bool, checkLeak bool) string {
	var password string = ""
	aPossibleValue := []string{}

	// génère le tableau de possibilités pour le mot de passe
	if lowLetter {
		aLowLetter := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
		aPossibleValue = append(aPossibleValue, aLowLetter...)
	}
	if bigLetter {
		aBigLetter := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
		aPossibleValue = append(aPossibleValue, aBigLetter...)
	}
	if number {
		aNumber := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
		aPossibleValue = append(aPossibleValue, aNumber...)
	}

	if specialCharacter {
		// TODO voir les vrais caractères utilisables
		aSpecialChar := []string{".", "!", ",", "@", ";", "?", "$"}
		aPossibleValue = append(aPossibleValue, aSpecialChar...)
	}

	// boucle qui recrée un mot de passe si il a leak

	// génère un mot de passe
	// petite ruse pour générer la taille aléatoire du mot de passe, on soustrait la longueur min du mot de passe pour générer un nombre puis on rajoute la longueur min
	tmpLen, _ := rand.Int(rand.Reader, big.NewInt(int64(maxLen)-int64(minLen)))
	lenPwd := tmpLen.Int64() + int64(minLen)

	for i := 0; i < int(lenPwd); i++ {

		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(aPossibleValue))))
		password += aPossibleValue[index.Int64()]
	}

	// vérifie si le mot de passe à leak
	if checkLeak {
		if checkPasswordLeak(password) {

		}
	}

	return password
}
func checkPasswordLeak(pwd string) bool {

	// hash le mot de passe
	passwordHash := HashString(pwd)

	// on récupère les hashs qui possèdent les mêmes 5 premiers caractère que notre hash
	strHashPossible, err := GetHtmlFromUrl("https://api.pwnedpasswords.com/range/" + passwordHash[:5])
	if err != nil {
		// si erreur lors de la récupération des infos sur le site haveibeenpwned
		return false
	} else {
		var passwordLeaked bool = false
		// on split les infos une première fois pour charger chaque possibilité
		tabPossibilities := strings.Split(strHashPossible, "\n")

		for _, elt := range tabPossibilities {
			// on resplite pour avoir le hash à tester ainsi que le nombre de fois ou il a leak si il a leak
			tabElt := strings.Split(elt, ":")
			if tabElt[0] == passwordHash[5:] {
				passwordLeaked = true
				break
			}
		}
		return passwordLeaked
	}
}

func HashString(str string) string {
	h := sha1.New()
	h.Write([]byte(str))

	hash := fmt.Sprintf("%X", h.Sum(nil))

	return hash

}

func GetHtmlFromUrl(url string) (string, error) {

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Erreur ", err)
		return "", errors.New("Problème de connexion internet ou erreur dans l'url")
	}

	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Erreur", err)
		return "", errors.New("Erreur lors de la lecture ")
	}
	return string(html), nil
}
