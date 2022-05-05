package CodingInterviewQuestions;

/**
 * GeeksForGeeks - https://practice.geeksforgeeks.org/problems/winner-of-an-election-where-votes-are-represented-as-candidate-names-1587115621/1
 * 
 * 
 */
import java.util.HashMap;

public class FindWinnerInElections {
	
	// Return the winner from the votes array of the strings
	public static String findWinner(String[] votes) {
		
		String winner = "";
		int winnerCount = 0;
	
		HashMap<String, Integer> candidatesList = new HashMap<String, Integer>();
		
		for (int i=0; i < votes.length; i++) {
							
			if (candidatesList.containsKey(votes[i])) {
				
				//System.out.println("Key is contained - " + votes[i]);
				int currentCount = candidatesList.get(votes[i]);
				currentCount = currentCount+1;
				
				candidatesList.put(votes[i], currentCount);
			}
			else {
				candidatesList.put(votes[i], 1);
			}
				
		}

		
		for( String candidateName: candidatesList.keySet()) {
			
			//System.out.println("Candidate name -> "+ candidateName  + ", Count ->"+ candidatesList.get(candidateName));
			
			if (candidatesList.get(candidateName) > winnerCount) {
				winner = candidateName;
				winnerCount = candidatesList.get(candidateName) ;
				
			}
			
		}
		
		return winner;
		
	}
	
	
	public static void main(String[] args){
		
		String[] votes = {"Anshul", "Ankit", "Anurag", "Anshul", "Ankit", "Ankit", "Anshul", "Anshul"};
		
		String winner = findWinner(votes);
		
		System.out.println(winner);
	}
 }
