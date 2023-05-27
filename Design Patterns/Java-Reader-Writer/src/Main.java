import java.io.*;
import java.lang.reflect.Array;
import java.nio.charset.Charset;
import java.nio.charset.StandardCharsets;
import java.util.Arrays;

public class Main {
	public static void main(String[] args) throws IOException {

		System.out.println("Hello world!");

		/*
		 * Reader class of java.io package is an abstract superclass that represents a stream of characters
		 *						 Reader class (abstract class)
		 * 							   |
		 * 							   |extends (up)
		 * 			   	-------------------------------------
		 * 				|				|					|
		 * 		BufferedReader		InputStreamReader		StringReader
		 * 								|
		 * 								| extends(up)
		 * 								|
		 * 							FileReader
		 *
		 *
		 *
		 */

		// creates a reader
		// Reader reader = new FileReader();

		/*
		* Few methods of reader
		* 		1. ready() -> checks if the reader is ready to be read.
		* 		2. read(char[] arr) -> reads the character from the stream and stores them in the specified array.
		* 		3. read(char[] arr, int start, int length) -> reads number of character equal to length from start and
		* 													stores them in arr
		* 		4. mark() -> marks the position in the stream.
		* 		5. reset() -> returns control to the point where mark was set.
		* 		6. skip() -> discards the specified number of characters from the stream.
		*
		 */

		char[] arr = new char[100];
		try {
			Reader reader = new FileReader("input.txt");
			System.out.println("Is there a stream : " + reader.ready());
			int num = reader.read(arr);
			System.out.println("Number of characters read : " + num);
			System.out.println("Data in the stream ");
			System.out.println(Arrays.toString(arr));
			reader.close();
		} catch (Exception e) {
			e.printStackTrace();
		}

		/*
		 * Writer class of java.io package is an abstract superclass that represents a stream of characters
		 *						 Writer class (abstract class)
		 * 								|
		 * 							    |extends (up)
		 * 			   	-------------------------------------
		 * 				|				|					|
		 * 		BufferedWriter		InputStreamWriter		StringWriter
		 * 								|
		 * 								| extends(up)
		 * 								|
		 * 							FileWriter
		 *
		 *
		 */

		//few methods which are related to writer.
		/*
		*	1. write(char[] arr) -> writes the characters from the specified array to the output stream
		* 	2. write(String data) -> writes the specified string to the writer.
		* 	3. append(char c) -> inserts a specified character to the current writer.
		* 	4. flush() -> forces to write all the data present in the writer to the specified destination.
		* 	5. close() -> closes the writer
		 */
		try {
			String data = "This is the data in the output file";
			Writer writer = new FileWriter("output.txt");
			writer.write(data);
			writer.close();
		} catch (Exception ex) {
			ex.printStackTrace();
		}

		/*
		*		Java InputStreamReader
		*	The input stream reader works with other streams as well.
		* 	It is also known as bridge between byte streams and character streams.
		* 	This is because InputStreamReader reads byte from the input stream as characters
		 */

		try {
			char[] a = new char[100];
			FileInputStream fs = new FileInputStream("input.txt");
			InputStreamReader isr = new InputStreamReader(fs);
			int num = isr.read(a);
			System.out.println("number of characters read : " + num);
		} catch (Exception ex) {
			ex.printStackTrace();
		}

		// getEncoding method can be used to get the type of encoding that is used to store the data in input stream.
		try {
			FileInputStream fs = new FileInputStream("input.txt");
			InputStreamReader isr = new InputStreamReader(fs); // this creates a input stream with default encoding

			InputStreamReader isr2 = new InputStreamReader(fs, StandardCharsets.UTF_8);

			System.out.println("encoding  for isr : " + isr.getEncoding());
			System.out.println("encoding for isr2 : " + isr2.getEncoding());
			char[] a = new char[100];
			System.out.println(isr.read(a));
			System.out.println(a);

			isr.close();
			isr2.close();

		} catch(Exception e) {
			e.printStackTrace();
		}

		/*
		 *	OutputStreamWriter
		 *
		 * 	Works with other output stream as well.
		 * 	It is also known as a bridge between byte stream and character streams.
		 * 	This is bc OutputStreamWriter convert its character to bytes.
		 */

		try {
			FileOutputStream fos = new FileOutputStream("output2.txt");
			OutputStreamWriter osw = new OutputStreamWriter(fos);

			// the above osw can also be created as below
			OutputStreamWriter osw2 = new OutputStreamWriter(fos, StandardCharsets.UTF_8);

			osw.write("This is to test output stream writer");
			System.out.println("Encoding : " + osw.getEncoding());
			System.out.println("Encoding 2 : " + osw2.getEncoding());
			osw.close();
			osw2.close();
		} catch (Exception e) {
			e.printStackTrace();
		}

		/*
		*		Java FileReader class -> to read data in characters
		 */

		try {
			FileReader fr = new FileReader("input.txt");
			char[] a = new char[100];
			System.out.println(fr.read(a));
			System.out.println("Data in file : " + Arrays.toString(a));
			// another approach can be to use the FileReader with specifying the charset
			fr.close();
		} catch (Exception e) {
			e.printStackTrace();
		}

		/*
		* 	Java FileWriter -> writes in chars to files
		*
		 */

		try {
			FileWriter fw = new FileWriter("output3.txt"); // encoding can also be provided here
			String d = "new string hehe";
			fw.write(d);
			fw.close();

		} catch (Exception ex) {
			ex.printStackTrace();
		}

		/*
		*				BufferedReader -> It can be used with other readers to read data in characters
		*
		* It extends the reader class. During the read operation a chunk of characters are read from disk and store in
		* an internal buffer, from this buffer characters are read individually.
		* Hence, the number of communication to the disk is reduced.
		 */

		try {
			FileReader fr =new FileReader("input.txt");
			BufferedReader br = new BufferedReader(fr);
			char[] a = new char[100];
			br.read(a);
			System.out.println("---" + Arrays.toString(a));
			br.close();
		} catch (Exception ex) {
			ex.printStackTrace();
		}

		//just testing can be ignored
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		String str = br.readLine();
		System.out.println(str);

		/*
		* Java BufferedWriterClass. -> can be used with other writers to write characters more efficiently.
		* During the write operation the characters are written to an internal buffer instead of the disk.
		* Once the buffer is filled or the writer is closed, the whole character is then written to the disk.
		 */

		try {
			String str2 = "New string to demonstrate the BufferedWriter class in java";
			FileWriter fw = new FileWriter("output4.txt");
			BufferedWriter bw = new BufferedWriter(fw);
			bw.write(str2);
			bw.close();
		} catch (Exception e) {
			e.printStackTrace();
		}

		//same as earlier flush function can also be used with this.

		/*
		* 								StringReader class (extends the Reader abstract class)
		* In StringReader class the string specified acts as source here characters are read individually.
		 */

		try {
			String st4 = "abc";
			StringReader sr = new StringReader(st4);
			char[] a = new char[100];
			System.out.println(sr.read(a));
			System.out.println(a);
		} catch(Exception ex) {
			ex.printStackTrace();
		}

		/*
		 * 								StringWriter class (extends the Writer abstract class)
		 * In Java StringBuffer is considered as mutable string.
		 *
		 */

		try {
			String st = "This is a new string from future";
			StringWriter sw = new StringWriter();
			sw.write(st);
			System.out.println(sw);
			sw.close();
		} catch (Exception ex) {
			ex.printStackTrace();
		}

		/*
		Using StringWriter with StringBuffer
		 */

		try {
			String st = "some string hehasdas asdal kdjaklsd a";
			StringWriter sw = new StringWriter();
			sw.write(st);
			StringBuffer sb = sw.getBuffer();
			System.out.println(sb.toString());
			sw.close();
		} catch (Exception ex) {
			ex.printStackTrace();
		}

		/*
		*	PrintWriter -> Converts primitive data(int, float, char etc...) into text format.
		* 	It then prints the formatted data to the writer.
		* 	It also doesn't throw any exceptions
		*
		* 	Here we need to use checkError() if we want to check for errors
		 */


		/*
		 *	Different ways to create a PrintWriter.
		 *	//creates a file writer
		 * 	1. Using writers :
		 * 		FileWriter fw = new FileWriter("file.txt");
		 * 		PrintWriter pw = new PrintWriter(fw);
		 *
		 * 	2. Using other output streams
		 * 		FileOutputStream fos = new FileOutputStream("output.txt");
		 * 		PrintWriter pw = new PrintWriter(fos, autoFlush); // autoflush is an optional param
		 *
		 *  3. Using filename
		 * 		PrintWriter pw = new PrintWriter(String fileName, boolean autoFlush)
		 * 								OR
		 *
		 * 		PrintWriter pw = new PrintWriter(String fileName, boolean autoFlush, Charset)
		 */

		try {
			String word= "my name is .sdklsdks";
			PrintWriter pw = new PrintWriter("output5.txt");
			pw.print(word);
			pw.close();
		} catch (Exception ex) {
			ex.printStackTrace();
		}

		//similarly printf can also be used.

	}
}