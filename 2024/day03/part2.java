import java.io.*;
import java.util.regex.Pattern;
import java.util.regex.Matcher;

public class Part2 {
    static int calcMul(String s) {
        int result = 0;
        Pattern patt = Pattern.compile("mul\\((\\b\\d{1,3}\\b),(\\b\\d{1,3}\\b)\\)");
        Matcher m = patt.matcher(s);
        while(m.find()) {
            int a = Integer.parseInt(m.group(1)); 
            int b = Integer.parseInt(m.group(2)); 
            result += a * b;
        }
        return result;
    }

    public static void main(String[] args) {
        long result = 0;
        try {
            boolean isActive = true;
            File f = new File("./input");
            BufferedReader br = new BufferedReader(new FileReader(f));
            String s;

            Pattern pattDont = Pattern.compile("don't\\(\\)");
            Pattern pattDo = Pattern.compile("do\\(\\)");

            while((s = br.readLine()) != null) {
                Matcher matDont = pattDont.matcher(s);
                Matcher matDo = pattDo.matcher(s);
                while(!s.isEmpty()) {
                    if(isActive) {
                        if(matDont.find()) {
                            result += calcMul(s.substring(0, matDont.start()));
                            s = s.substring(matDont.end());
                            matDont = pattDont.matcher(s);
                            matDo = pattDo.matcher(s);
                            isActive = false;
                        } else {
                            result += calcMul(s);
                            s = "";
                        }
                    } else {
                        if(matDo.find()) {
                            s = s.substring(matDo.end());
                            matDont = pattDont.matcher(s);
                            matDo = pattDo.matcher(s);
                            isActive = true;
                        } else {
                            s = "";
                            matDont = pattDont.matcher(s);
                            matDo = pattDo.matcher(s);
                        }
                    }
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

