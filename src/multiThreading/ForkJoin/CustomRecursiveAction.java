package multiThreading.ForkJoin;

/**
 * https://www.baeldung.com/java-fork-join
 * 
 * 
 */
import java.util.*;
import java.util.concurrent.ForkJoinTask;


public class CustomRecursiveAction {
	
	private String workload = "";
    private static final int THRESHOLD = 4;

    //private static Logger logger = 
    //  Logger.getAnonymousLogger();

    public CustomRecursiveAction(String workload) {
        this.workload = workload;
    }

    @Override
    protected void compute() {
        if (workload.length() > THRESHOLD) {
            ForkJoinTask.invokeAll(createSubtasks());
        } else {
           processing(workload);
        }
    }

    private List<CustomRecursiveAction> createSubtasks() {
        List<CustomRecursiveAction> subtasks = new ArrayList<>();

        String partOne = workload.substring(0, workload.length() / 2);
        String partTwo = workload.substring(workload.length() / 2, workload.length());

        subtasks.add(new CustomRecursiveAction(partOne));
        subtasks.add(new CustomRecursiveAction(partTwo));

        return subtasks;
    }

    private void processing(String work) {
        String result = work.toUpperCase();
        //logger.info("This result - (" + result + ") - was processed by " 
        //  + Thread.currentThread().getName());
    }

}
