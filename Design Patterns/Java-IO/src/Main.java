import java.io.*;
import java.nio.file.Files;
import java.nio.file.Paths;

// Reference : https://www.programiz.com/java-programming

public class Main {
	public static void main(String[] args) {
		// input stream -> read from a data source.
		// output stream -> write data to data destination.
		System.out.println("Hello world!"); // here System.out is output stream.

		/*
		 * Types of streams :
		 * 		1. Byte Stream.
		 * 		2. Character Stream.
		 * <p>
		 *
		 * 	Byte Stream :
		 * 		It is used to read and write a single byte(8 bits) of data
		 * 		All byte stream classes are derived from base abstract classes
		 * 		called InputStream, and OutputStream.
		 * <p>
		 *
		 * 	Character Stream :
		 * 		It is used to write a single character of data.
		 * 		All character stream classes are derived from Reader and Writer.
		 */

		/*
		 * 	InputStream class :  of java.io package is an abstract superclass that
		 *	represents input stream of bytes.
		 *
		 * 	As InputStream is an abstract class, so it is not useful by itself.
		 * 	Its subclasses are :
		 * 			1. FileInputStream
		 * 			2. ByteArrayInputStream
		 * 			3. ObjectInputStream
		 */

		/*
		 * Commonly used methods of InputStream
		 * 1. read() -> reads one byte data from input stream.
		 * 2. read(byte[] array) -> reads bytes from input stream and stores them in byte array.
		 * 3. available() -> return number of bytes available in input stream
		 * 4. mark() -> marks the position in the input stream upto which data has been read.
		 * 5. reset() -> returns the control to the point in the stream where the mark was set.
		 * 6. markSupported() -> checks if mark() and reset() are supported in the stream.
		 * 7. skips() -> skips and discards specified number of bytes from the stream.
		 * 8. close() -> closes the stream.
		 */

		//example using FileInputStream.
		byte[] arr = new byte[100];
		try {
			InputStream ip = Files.newInputStream(Paths.get("test.txt"));
			System.out.println("Available bytes : " + ip.available());
			//read byte from input stream
			System.out.println(ip.read(arr));

			//to String
			String data = new String(arr);
			System.out.println(data);
			ip.close();
		} catch (Exception e) {
			e.printStackTrace();
		}
		/*
		 * OutputStream :
		 * 		class of java.io package is an abstract superclass that represents output stream of bytes.
		 * 	Subclasses of OutputStream
		 * 	1. FileOutputStream
		 * 	2. ByteArrayOutputStream
		 * 	3. ObjectOutputStream
		 */
		//example of creating an output stream
		String data = "this is a temporary data.";
		try {
			OutputStream op = Files.newOutputStream(Paths.get("output.txt"));
			byte[] d = data.getBytes();
			op.write(d);
			System.out.println("Data is written to the file");
			op.close();
		} catch (Exception e) {
			e.printStackTrace();
		}

		//example of using skip
		try {
			InputStream ip = Files.newInputStream(Paths.get("test.txt"));
			System.out.println("Available bytes : " + ip.available());
			//read byte from input stream
			byte[] byteArr = new byte[100];
			System.out.println(ip.skip(5));
			System.out.println("input after skipping 5 bytes");
			int i = ip.read();
			while(i != -1) {
				System.out.println((char)i);
				i = ip.read();
			}

			//to String
			ip.close();
		} catch (Exception e) {
			e.printStackTrace();
		}

		/*
		 * the flush() method in JavaOutputStream
		 * to clear the output stream we use flush
		 * It forces the output stream to write all the
		 * data to the destination.
		 */

		String data2 = "this is a temp data to test flush";
		try {
			OutputStream fp = Files.newOutputStream(Paths.get("output2.txt"));

			fp.write(data2.getBytes());//using the write method
			//using flush method
			fp.flush();
			fp.close();
		} catch (Exception e) {
			e.printStackTrace();
		}

		/*
		 * ByteArrayInputStream :
		 * 		Class of java.io package can be used to read array of input bytes.
		 * 		It extends the InputStream.

		 * 	-----creating a byte array-----
		 * 		ByteArrayInputStream br = new ByteArrayInputStream(byte[] arr);

		 * 	-----creating a byte array to read only a portion of array-----
		 *		ByteArrayInputStream inbr = new ByteArrayInputStream(byte[] arr, int start, int length);

		 *	Common methods:
		 *		1. read () : reads a single byte from the array present in the input stream.
		 *		2. read(byte[] arr) : reads byte from the specified input stream and stores them in the arr.
		 *		3. read(byte[] arr, int start, int length) : reads number of byte equal to the length from start.
		 */

		//example of byte array input stream
		byte[] a = {1, 2, 3, 4};
		try  {
			InputStream in = new ByteArrayInputStream(a);
			System.out.println("Reading bytes from the array");
			for (int i=0; i<a.length; i++) {
				int databyte = in.read();
				System.out.print(databyte + " , ");
			}
			System.out.println();
			in.close();
		} catch (Exception e) {
			e.printStackTrace();
		}

		// to get the number of bytes available, use the in.available() to get how much bytes are available
		// in.skip(int val); to val bytes from the stream.

		/*
		 *	ByteArrayOutputStream :
		 * 		use to write array of output data in bytes
		 *
		 * 	Creation :
		 * 		ByteArrayOutputStream br = new ByteOutputStream();
		 * 	can also specify the size
		 * 		ByteArrayOutputStream br = new ByteOutputStream(int size);
		 *
		 * Some common methods:
		 * 		1. write(byte[] arr) : write the bytes from the arr  to the output stream.
		 * 		2. write(int byte) : writes specified number of bytes to the output stream.
		 * 		3. write(byte[] arr, int start, int length) : write specified length fo bytes from the array to the
		 * 			output stream.
		 * 		4. writeTo(ByteArrayOutputStream br) : write all the bytes from current output stream to br.
		 */

		//example of byte output stream
		String str = "this is a line of string to demonstrate the byteoutputarraystream";
		try {
			ByteArrayOutputStream br = new ByteArrayOutputStream();
			byte[] array = str.getBytes();

			//write data to the
			br.write(array);
			String strOutputStream = br.toString();
			System.out.println("Data  :"  + strOutputStream);
		} catch(Exception ex) {
			ex.printStackTrace();
		}

		/*
		 *	to access data from output stream
		 * 	Two options :
		 * 			1. toByteArray()
		 * 			2  toString()
		 */

		String data3 = "this is data";
		try {
			ByteArrayOutputStream br = new ByteArrayOutputStream();
			br.write(data3.getBytes());
			byte[] data3Arr = br.toByteArray();
			for (int i=0; i<data3Arr.length; i++) {
				System.out.print((char)data3Arr[i] + ", ");
			}
			String data3string = br.toString();
			System.out.println(data3string);
			br.close();
		} catch (Exception e) {
			e.printStackTrace();
		}

		/*
		Some other useful methods:
				1. size() -> returns the size of the array.
				2. flush() -> clears the output stream.
		 */

      //---------------------------------------------------------------------------
		/*
		 * ObjectInputStream :
		 * 		class of java.io package.
		 * 		It is used to read objects that were previously written by ObjectInputStream
		 * It extends InputStream.

		 * ObjectOutputStream is used to write objects to ObjectInputStream
		 * It extends the OutputStream.
		 *
		 * ObjectOutputStream encodes the Java Objects using class name and object values.
		 * hence generates corresponding streams. This process is known as serialization.
		 *
		 * Those converted streams can be stored in files and can be transferred over networks.
		 * ObjectOutputStream only writes objects that implement Serializable interface.
		 *
		 *
		 * Methods :
		 * 		1. write() -> writes a byte of data to the output stream.
		 * 		2. writeBoolean() -> writes data in boolean form.
		 * 		3. writeChar() -> writes data in char form.
		 * 		4. writeInt() -> write data int form.
		 * 		5. writeObject() -> write object to output stream.
		 *
		 */

		//example for ObjectOutputStream
		int dataInt = 1;
		String strData = "This is some example.";
		try {
			OutputStream file = Files.newOutputStream(Paths.get("test.txt"));

			// create an object output stream
			ObjectOutputStream os = new ObjectOutputStream(file);
			os.writeInt(dataInt);
			os.writeObject(strData);

			// reads data from object input stream
			InputStream file2 = Files.newInputStream(Paths.get("test.txt"));
			ObjectInputStream inOs = new ObjectInputStream(file2);
			System.out.println("Integer data : " + inOs.readInt());
			System.out.println("Read Object : " + inOs.readObject());

			inOs.close();
			os.close();
		}catch (Exception e) {
			e.printStackTrace();
		}

		//example two using a class which implements serializable.
		Dog dog = new Dog("dog1", "shiba inu");
		try {
			OutputStream os = Files.newOutputStream(Paths.get("file.txt"));
			ObjectOutputStream oos = new ObjectOutputStream(os);

			oos.writeObject(dog);

			//read the object from file.
			InputStream is = Files.newInputStream(Paths.get("file.txt"));
			ObjectInputStream ois = new ObjectInputStream(is);
			Dog d = (Dog) ois.readObject();
			System.out.println("Dog name : " + d.name);
			System.out.println("Breed : " + d.breed);
			os.close();
			oos.close();
		} catch (Exception e) {
			e.printStackTrace();
		}

		/**
		 * ObjectInputStream :
		 * 		It is mainly used to read written data from ObjectOutputStream
		 * 		Methods :
		 * 			1. read()
		 * 			2. readBoolean()
		 * 			3. readInt()
		 * 			4. readObject()
		 * 			5. readChar()
		 *
		 * 	Other methods :
		 * 		1. available() -> returns number of bytes available in input stream
		 * 		2. mark() -> marks the position in input stream upto which data has been read.
		 * 		3. reset() -> returns the control to the point in the input stream where mark was set.
		 * 		4. skipBytes() -> skips and discards bytes from the specified stream.
		 * 		5. close() -> closes the stream.
		 *
		 */


		/*
		* 							BufferedStreamClass
		* 							/					\
		* 			BufferedInputStream					BufferedOutputStream
		* 		Maintains a input buffer
		* 		of 8192 bytes.
		*
		* During the read operation in BufferedInputStream, a chunk of bytes is read from the disk
		* and stored in the internal buffer. And from the internal buffer bytes are read individually.
		* Hence the number of communication to the disk is reduced.
		* This is the reason why Buffered Streams are faster.
		*
		* Few Methods :
		* 		1. read() -> read a single byte from the input stream.
		* 		2. read(byte[] arr) -> reads bytes from the input stream and stores in the byte arr.
		* 		3. read(byte[] arr, int start, int length) -> reads number of byte equal to the length.
		*
		 */

		try {
			FileInputStream fs = new FileInputStream("input.txt");
			BufferedInputStream bfs = new BufferedInputStream(fs);

			int i = bfs.read(); // read the first byte from the stream
			while(i != -1) {
				System.out.println((char)i);
				i = bfs.read();
			}
			bfs.close();
		} catch (Exception e) {
			e.printStackTrace();
		}

		// it also has available method to check number of bytes available in input stream.
		// skip discards the specified number of bytes.

		/*
		* BufferedOutputStream:
		* 		maintains internal buffer of 8192 bytes.
		*
		* Methods:
		* 		1. write() -> writes a single byte to the internal buffer of output stream.
		* 		2. write(byte[] arr) -> writes bytes from specified array to the output stream.
		* 		3. write(byte[] arr, int start, int length) -> writes number of bytes equal to the length.
		*
		 */

		String str1 = "this is a line of text to demonstrate buffered output stream";
		try {
			FileOutputStream fs = new FileOutputStream("output.txt");
			BufferedOutputStream bfs = new BufferedOutputStream(fs);
			byte[] arr6 = str1.getBytes();
			bfs.write(arr6);
			bfs.close();
		} catch (Exception e) {
			e.printStackTrace();
		}

		// flush() method to clear the output stream to write all data present in the buffer to the destination.

		/*
		* 								PrintStream
		* This is used to write output data in commonly readable format which is text.
		* It extends the abstract class OutputStream.
		* PrintStream does not throw any exceptions or anything, we need to check using
		* checkError() method to see if there is any error.
		*
		* PrintStream has a feature of auto flushing which means force the data to write all the data
		* to the destination, under following conditions.
		* 1. if newline character is present. \n
		* 2. println() method invoked.
		* 3. if an array of bytes is written in print stream
		*
		*
		 */

		//example :
		try {
			String val = "this is some junk using print stream";
			PrintStream ps = new PrintStream("output.txt"); // there either we can pass file name or FileOutputStream
			ps.print(val);
			ps.close();
		}catch (Exception e) {
			e.printStackTrace();
		}
	}
}

class Dog implements Serializable {
	String name;
	String breed;
	public Dog(String name, String breed) {
		this.name = name;
		this.breed = breed;
	}
}