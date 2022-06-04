package CodingInterviewQuestions;

/**
 * Asked in Vonage Coding Round, 4thJune2022.
 */

/*import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.bind.annotation.RestController;

import java.time.ZonedDateTime;
import java.util.HashMap;

@RestController
class HealthCheckController {

    // Your solution

    @GetMapping(value = "/healthcheck")
    public ResponseEntity healthcheck(@RequestParam(value="format", required=false) String format) {

        HashMap<String, String> responseBody = new HashMap<>();

        responseBody.put("status", "OK");

        if (format != null && !format.isEmpty()) {

            switch(format) {

                case "short":
                    return ResponseEntity.ok().body(responseBody);

                case "full":
                    responseBody.put("currentTime", String.valueOf(ZonedDateTime.now()));
                    return ResponseEntity.ok().body(responseBody);

                default:
                    //nothing
            }
        }

        return ResponseEntity.status(400).build();
    }

    @PutMapping(value = "/healthcheck")
    public ResponseEntity healthcheckPut() {
        return ResponseEntity.status(405).build();
    }

    @PostMapping(value = "/healthcheck")
    public ResponseEntity healthcheckPost() {
        return ResponseEntity.status(405).build();
    }


    @DeleteMapping(value = "/healthcheck")
    public ResponseEntity healthcheckDelete() {
        return ResponseEntity.status(405).build();
    }

}*/
