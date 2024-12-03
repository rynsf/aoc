import java.io.*;
import java.util.regex.Pattern;
import java.util.regex.Matcher;

public class Part1 {
    public static void main(String[] args) {
        long result = 0;
        try {
            File f = new File("./input");
            BufferedReader br = new BufferedReader(new FileReader(f));
            String s;

            Pattern patt = Pattern.compile("mul\\((\\b\\d{1,3}\\b),(\\b\\d{1,3}\\b)\\)");
            Matcher m;
            while((s = br.readLine()) != null) {
                m = patt.matcher(s);
                while(m.find()) {
                    int a = Integer.parseInt(m.group(1)); 
                    int b = Integer.parseInt(m.group(2)); 
                    result += a * b;
                }
            }
        } catch(FileNotFoundException e) {
            System.err.println("file not found");
        } catch(IOException e) {
            System.err.println("IO exception");
        }

        System.out.println(result);
    }
}
