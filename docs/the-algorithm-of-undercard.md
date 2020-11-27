# Undercard

Undercard is designed to solve the security problem of random number generation algorithm.  

Random numbers in the game are very critical, it can not only increase the fun of the game, in many cases also play a role in determining the game win or lose, and that has a direct impact on the fairness of the game.  

__The first__  
As a random number in the game, validity is short-term. Therefore, as a random number generation algorithm, it should be able to guarantee the security of random numbers during the validity period of random numbers, and provide verification capability after the expiration of the validity period of random numbers.  

__Sencond__  
The destructiveness of random numbers being leaked or controlled during the validity period is enormous and is therefore the core problem that the algorithm wants to solve.  
The solution to random number disclosure is to use encryption;  
There are two main extremes in which random numbers are controlled:  
    - Servers that produce random numbers are controlled;  
    - There is only one honest player, and everyone else is colluding.  
Because of this, the primary purpose of the validation mechanism is to verify that random numbers are not controlled.  

__The last__  
The current method of generating random numbers does not necessarily guarantee true randomness. The traditional random number generation algorithm itself is a pseudo-random algorithm, or increases randomness by the physical amount of some servers themselves. But this needs to be done in a credible environment. In a decentralized environment, this method makes the verification of random numbers very difficult. It is difficult to verify that these parameters in the process of generating random numbers are real and not controlled by the server.  
Therefore, ensuring the randomness and verifiability of random numbers is the ultimate design goal of this algorithm.  
