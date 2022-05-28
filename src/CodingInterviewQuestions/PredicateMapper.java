package CodingInterviewQuestions;

/**
 * Asked in Twillo Coding Round, 28thMay2022.
 *
 * To-Be-Done.
 *
 * Similar LeetCode Question -
 * - Reference Link - https://docs.oracle.com/javase/8/docs/api/java/util/function/Predicate.html
 */

import java.util.Arrays;
import java.util.List;
import java.util.Scanner;

import java.util.function.*;
/*
 * Create the Filter and Mapper classes here.
 */

/*class Filter {

    // I don't know what is Predicate class ( or internal implementation). Hence unable to solve this. Otherwise, i can solve this.
    static Predicate<String> nameStartingWithPrefix(String prefix) {

        //Predicate obj = new Predicate();

        return null;
    }
}

class Mapper {

    // I don't know what is Predicate class. Hence this is the problem. Hence unable to solve this. Otherwise, i can solve this.
    static Function<String, CharactersCount> getDistinctCharactersCount() {

        //stream.

        // input string
        //CharactersCount count = new CharactersCount(name, distinctCharacterCount)

        // We can use HashMpa to build the characterscount in the string.
        // I don't know how to use java stream, predicate, function. hence this is the proble,

        return null;
    }
}
class CharactersCount {
    private final String name;
    private final Integer distinctCharacterCount;

    public CharactersCount(String name, Integer distinctCharacterCount) {
        this.name = name;
        this.distinctCharacterCount = distinctCharacterCount;
    }

    @Override
    public String toString() {
        return "\"" + this.name + "\" has " + this.distinctCharacterCount + " distinct characters.";
    }
}*/

public class PredicateMapper {
    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) {
        List<String> names = Arrays.asList(
                "aaryanna",
                "aayanna",
                "airianna",
                "alassandra",
                "allanna",
                "allannah",
                "allessandra",
                "allianna",
                "allyanna",
                "anastaisa",
                "anastashia",
                "anastasia",
                "annabella",
                "annabelle",
                "annebelle"
        );

        System.out.println("Test1");

        /*names.stream()
                .filter(Filter.nameStartingWithPrefix(scanner.nextLine()))
                .map(Mapper.getDistinctCharactersCount())
                .forEachOrdered(System.out::println);*/
    }
}
