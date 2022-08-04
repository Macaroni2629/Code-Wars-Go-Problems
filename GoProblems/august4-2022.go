// https://www.codewars.com/kata/59ca8246d751df55cc00014c/train/go

// A hero is on his way to the castle to complete his mission. However, he's been told that the castle is surrounded with a couple of powerful dragons! each dragon takes 2 bullets to be defeated, our hero has no idea how many bullets he should carry.. Assuming he's gonna grab a specific given number of bullets and move forward to fight another specific given number of dragons, will he survive?

// Return True if yes, False otherwise :)

//My solution

package kata

func Hero(bullets, dragons int) bool {
  dragons = dragons * 2
  if bullets >= dragons {
    return true
  }
  return false
}

// Other solutions

package kata

func Hero(bullets, dragons int) bool {
  return bullets >= 2*dragons
}


