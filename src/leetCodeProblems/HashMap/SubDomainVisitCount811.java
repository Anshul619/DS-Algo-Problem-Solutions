package leetCodeProblems.HashSearch;

/**
 * LeetCode - https://leetcode.com/problems/subdomain-visit-count/solution/
 *
 * TimeComplexity - O(N)
 * SpaceComplexity - O(N)
 */

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;

public class SubDomainVisitCount811 {

    HashMap<String, Integer> subDomainsFrequencyMap;

    private void incrementCount(String domain, Integer counter) {

        if (subDomainsFrequencyMap.containsKey(domain)) {
            int count = subDomainsFrequencyMap.get(domain);
            count += counter;
            subDomainsFrequencyMap.put(domain, count);
        }
        else {
            subDomainsFrequencyMap.put(domain, counter);
        }
    }
    public List<String> subdomainVisits(String[] cpdomains) {

        List<String> output = new ArrayList<>();
        subDomainsFrequencyMap = new HashMap<>();

        for (int i=0; i < cpdomains.length; i++) {

            String[] splitArray = cpdomains[i].split(" ");

            int currentCounter = Integer.parseInt(splitArray[0]);
            //System.out.println(domainSplit[1]);
            String[] subdomains = splitArray[1].split("\\.");

            //System.out.println(Arrays.toString(subdomains));

            if (subdomains.length > 0) {

                String currentDomain = subdomains[subdomains.length-1];
                incrementCount(currentDomain, currentCounter);

                for (int j=subdomains.length-2; j >= 0; j--) {
                    currentDomain = subdomains[j] + "." + currentDomain;
                    incrementCount(currentDomain, currentCounter);
                }
            }
            else {
                incrementCount(splitArray[1], currentCounter);
            }
        }

        for (String domain: subDomainsFrequencyMap.keySet()) {
            output.add(subDomainsFrequencyMap.get(domain)+ " " + domain);
        }

        return output;
    }

    public static void main(String[] args){

        String[] input = {"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"};

        SubDomainVisitCount811 obj = new SubDomainVisitCount811();
        System.out.println(obj.subdomainVisits(input));

    }
}
