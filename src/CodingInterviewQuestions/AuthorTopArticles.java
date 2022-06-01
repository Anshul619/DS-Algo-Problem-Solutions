import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.*;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.TimeUnit;
import java.time.Duration;
import java.util.concurrent.TimeoutException;

/**
 * Asked in Instawork, 1stJune2022
 *
 * To-Be-Completed
 */

public class AuthorTopArticles {

    static class Article {
        String articleName;
        int noOfComments;
    }

    int totalPages = 0;

    static PriorityQueue<Article> queue;

    private static String callAuthorArticlesAPI(String username, int pageNumber) throws ExecutionException, InterruptedException, TimeoutException {

        String url = "https://jsonmock.hackerrank.com/api/articles?author=" + username+"&page="+pageNumber;

        System.out.println(url);
        HttpClient httpClient = HttpClient.newBuilder()
                .version(HttpClient.Version.HTTP_2)
                .connectTimeout(Duration.ofSeconds(10))
                .build();

        HttpRequest request =
                HttpRequest.newBuilder()
                        .GET()
                        .uri(URI.create(url))
                        .build();

        CompletableFuture<HttpResponse<String>> response =
                httpClient.sendAsync(request, HttpResponse.BodyHandlers.ofString());

        String result = response.thenApply(HttpResponse::body).get(5, TimeUnit.SECONDS);

        System.out.println(result);

        //JSONParser parser = new JSONParser();
        //JSONObject json = (JSONObject) parser.parse(result);

        return result;
    }

    /*private static void pushArticlesInQueue(JSONArray articles) {

        for (JSONObject article: articles) {

            Article obj = new Article();

            if (article['title'] != null) {
                obj.articleName = article['title'];
            }
            else {
                obj.articleName = article['story_title'];
            }

            if (article['num_comments'] == null) {
                obj.noOfComments = 0;
            }

            queue.add(obj);
        }

    }*/

    static class PriorityComparator implements Comparator<Article> {

        public int compare(Article article1, Article article2) {

            return article2.noOfComments - article1.noOfComments;
        }
    }
    public static List<String> topArticles(String username, int limit) throws ExecutionException, InterruptedException, TimeoutException {

        if (limit == 0) {
            return new ArrayList<>();
        }
        /*JSONObject result = callAuthorArticlesAPI(username, 1);

        int total = result['total'];
        int totalPages = result['total_pages'];
        JSONObject articles = result[data];*/

        int total = 10;
        queue = new PriorityQueue<>(total, new PriorityComparator());

        return new ArrayList<>();
    }

    public static void main(String[] args) throws ExecutionException, InterruptedException, TimeoutException {
        System.out.println("Hello world!");

        AuthorTopArticles.topArticles("olalonde", 5);
    }
}