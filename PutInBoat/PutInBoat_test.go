// Kryssing 1 = hs og kylling "[rev korn ---v\ \hs kylling/_________________/ø---]"
// Kryssing 2 = hs drar tilbake alene "[korn rev ---v\ \hs/_________________/ø--- kylling]"
// Kryssing 3 = hs og rev i båt "[korn ---v\ \hs rev/_________________/ø--- kylling]"
// Kryssing 4 = hs tar med kylling tilbake "[korn ---v\ \hs kylling/_________________/ø--- rev]"
// Kryssing 5 = hs legger fra kylling og tar med korn "[kylling ---v\ \hs korn/_________________/ø--- rev]"
// Kryssing 6 = hs drar alene tilbake "[kylling---v\ \hs/_________________/ø---rev korn]"
// Kryssing 7 = hs og kylling "[---v\ \hs kylling/_________________/ø---hs kylling rev korn]"

// liste West har hs, rev, korn, kylling og skal over til liste East med Boat
// Så lenge det er noen items i west, kjør kode
// Når east har 4 items, så er koden ferdig
// Kjøre en algoritme som går gjennom west og flytter items til east
// Sett opp if-setninger for hva som er i båten samtidig
// rev > kylling > korn
// func loop_through_west_and_add_to_boat() {}

package PutInBoat

import "testing"

func TestInsertHS(t *testing.T) {
	wanted := "HS is in boat"
	got := "fox"
	if got != wanted {
		t.Errorf("Feil, fikk #{got}, ønsket #{wanted}")
	}
}

//func TestBoat(t *testing.T) {
//fmt.Println("Starting")
//east:= []string{}
//west:= []string{"korn", "hs", "kylling", "rev"}
//boat:= []string{}

//func bring_to_other_side() {
//	if (east.length == 4) {
//		fmt.print("done")
//	}
//}
//}
