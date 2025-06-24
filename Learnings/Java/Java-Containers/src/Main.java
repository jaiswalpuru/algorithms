import java.io.*;
import java.nio.file.Files;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Stack;

public class Main {
	public static void main(String[] args) throws IOException {
		System.out.println("Welcome to containers....");

		/*
		*    				-------------- Java Containers ---------------
		*
		* 	Containers mainly include Collections and Map.
		*
		* 	Collections are used to store objects, while Map is used to store pair of objects.
		*						(<<i>> = interface)
		* 									<<i>>
		*							    (Collection)
		* 							/		  |       \
		* 						Set			List	Queue
		* 						<<i>>		<<i>>	 <<i>>

		* 	HashSet, LinkedHashSet, and SortedSet<<i>> implements -----> SET <<i>>
		*	TreeSet implements SortedSet<<i>>

		*	ArrayList, Vector, and LinkedList implements --------> LIST <<i>>

		* 	LinkedList and PriorityQueue implements ---------> QUEUE <<i>>

		*
		* 	SET:
		* 		TreeSet -> Implements RedBlack trees for storing objects in an ordered fashion, search is not as efficient
		* 					Search : O(logn), cam be used to retrieve range of elements efficiently.
		*
		* 		HashSet -> Based on hashtable implementation. Items are not stored in orderly manner
		*
		* 		LinkedHashSet -> Lookup efficiency same as HashSet, and internally uses DLL to maintain the order.
		*
		* 	LIST:
		* 		ArrayList -> Implemented based on dynamic array, and supports random access.
		*
		* 		Vector -> Similar to array list and also thread safe, therefore it is slower than ArrayList
		*		(Stack class extend vector)
		*
		* 		`LinkedList -> based on the implementation of doubly linked list. Accessing the element is sequential
		* 					but insertion and deletion are quicker.
		* 					NOTE : Only linked list can be used as stack, queue, bidirectionalQueue
		*
		*`	QUEUE:
		* 		LinkedList -> can be used to implement two way queue
		*
		* 		PriorityQueue -> Based on the heap data structure.
		*
		*
		*
		* 						********** MAP<<i>> **************
		*	SortedMap <<i>>, HashTable, LinkedHashMap, HashMap implements MAP interface.
		* 	TreeMap implements SortedMap interface.
		*
		* 	TreeMap : Implemented using RedBlack trees.
		*
		* 	HashMap : Implemented using hash table.
		*
		* 	HashTable : Similar to HashMap but it is thread safe(so will be slower than HashMap). It is legacy class
		* 				and should not be used.
		* 				Instead, use ConcurrentHashMap to support thread safety.
		*
		* 	LinkedHashMap: Use a DoublyLinkedList to maintain the order of elements, in insertions order Least recently
		* 				used. (LRU)
		*
		 */

		/*								Design Patterns in Containers
		*
		* 	Iterator Pattern :
		* 		Collection inherits iterable<<i>> in which the iterator method is there.
		* 		List -> produces ListIterator which inherits Iterator
		*
		 */

		List<Integer> list = new ArrayList<>();
		list.add(1);
		list.add(123);
		for (Integer val : list) System.out.println(val);
		// or using streams
		list.forEach(System.out::println);


		/*
		*					Adapter Pattern
		* 	java.util.Arrays#asList() can convert array type to List type.
		* 	@SafeVarargs
		* 	public static <T> List<T> asList(T... a)
		* 	The parameters of as list are generic variable length params.
		*
		* 	Arrays of primitive cannot be used in asList, only of packaged types, eg shown below.
		 */


		Integer[] arr = {1, 2, 3, 4, 5};
		List l = Arrays.asList(arr);
		l.forEach(System.out::println);


		/*
		*											ArrayList
		* 	It is implemented based on array, it supports fast random access.
		*
		* 	public class ArrayList<E> extends AbstractList<E> implements List<E>, RandomAccess, Cloneable, java.io.Serializable
		*
		* 	The default size of array is 10.
		*
		* 	private static final int DEFAULT_CAPACITY = 10;
		*
		* 	While adding elements, it uses function ensureCapacityInternal() to ensure there is enough space in the array
		* 	if there is not then it expands the array using the grow() function.
		*
		* 	 			The size of the new capacity = oldCap + (oldCap >> 1);
		*	So the new capacity is 1.5 times the old capacity.
		* 	The expansion operation needs to call Arrays.copyOf() to copy the entire original array to the new array.
		*	This operation is very expensive, so it is best to specify the capacity when creating the array list.
		*
		*
		* 							DELETE
		* 	You need to call System.arrayCopy() to copy the elements after index+1, to the index position.
		* 	The time complexity of the operation is O(N)
		* 	Cost of deleting elements in ArrayList is very high.
		*
		* 						SERIALIZATION
		*
		*	ArrayList is implemented based on arrays and has the feature of dynamic expansion, so there might be free space
		* 	which do not need to be serialized them all.
		* 	The array elementData that holds the element is decorated as transient, which declares that array will not be
		* 	serialized by default.
		*
		* 	transient Object[] elementData; // non -private to simplify nested class access.
		*
		* 	Array list implements readObject() and writeObject() to control the part filled with elements.
		*
		* 	When serializing, you need to use the writeObject() of ObjectOutputStream to convert the object in output
		* 	byte stream, and output it, the writeObject() will reflectively call the Objects writeObject() to achieve
		* 	serialization, when there is writeObject() in the incoming object.
		*
		* 	The deserialization uses the readObject() method of ObjectInputStream, the principal is similar.
		*
		*
		*	File f = new File("input.txt");
		*	ArrayList alist = new ArrayList();
		*	ObjectOutputStream os = new ObjectOutputStream(Files.newOutputStream(f.toPath()));
		*	os.writeObject(alist);
		*/

		/*
		*					FAIL FAST
		* 	modCount is used to record the number of times the Array structure has changed. Structural changes
		* 	are add, delete, at least one element, or resize the internal array.
		* 	When performing such operations such as serialization or iteration, it is necessary to compare whether,
		* 	the modCount has changed before or after the operation, and if there is a change in modCount then
		* 	ConcurrentModificationException is thrown.
		*
		 */


		//---------------------------------------------------------------------------------------------------------


		/*
		*								VECTOR
		*
		* 	1. Synchronization :
		* 		It is implemented similar to ArrayList but using the synchronized keyword.
		* 		(synchronized is used so that there to handle ConcurrentModification)
		*
		* 	2. Expansion:
		*		The constructor of a vector can pass in the capacity/increment parameter, and its function is to increase
		* 		the cap by capIncrement during cap expansion. If the value of increment param is <=0 then the cap will
		* 		be doubled each time during expansion
		*
		*
		 */


		//---------------------------------------------------------------------------------------

		/*
		*					Comparison ArrayList vs Vector
		*
		* 	• Vector is synchronized so overhead is more than that of ArrayList, and the access speed is lower.
		* 	  It is better to use ArrayList instead of vector as synchronization operation can be controlled by the
		* 	  programmer.
		*
		* 	• Vector request 2 times its size each time when it is expanded, while ArrayList is 1.5 times.
		*
		*
		 */

		/*							Alternatives
		 *
		 *		You can use Collections.synchronizedList() to get thread safety.
		 * 		List<String> list = new ArrayList<>();
		 * 		List<String> synList = Collections.synchronizedList(list);
		 *
		 * 		You can also use CopyOnWriteArrayList class under the concurrent package.
		 * 		List<String> lst = new CopyOnWriteArrayList<>();
		 */

		/*
		*						CopyOnWriteArrayList
		*
		*		Read and Write separation.
		*
		* 		The write operation is performed on a copy array, and the read operation is performed on original array.
		* 		The read and write effect are separated and do not affect each other.
		*
		* 		After the write operation the original array needs to point to the new array.
		*
		* 		CopyOnWriteArrayList has its own flaws even though read is fast:
		*
		* 				1. Memory usage: Each time a write operation is being performed, a new array needs to be created.
		* 								so memory usage is twice the original.
		*
		* 				2. Data Inconsistency: The read operation cannot read the real time data, bc the data of write
		* 								operation has not been synchronized to the read array.
		*
		*
		* 		So CopyOnWriteArrayList is not suitable for memory sensitive and high real-time applications.
		*
		 */

		//---------------------------------------------------------------------------------------

		/*
		*										Linked List
		*
		* 	Based on the implementation of doubly linked list, use Node to store the node info of linked list.
		*
		* 					private static class Node<E> {
		* 						E item;
		* 						NodeL<E> next;
		* 						Node<E> prev;
		* 					}
		*
		* 	Each linked list stores the first and last pointers.
		*
		* 		transient Node<E> first;
		* 		transient Node<E> last;
		 */

		/*
		*							Linked List VS ArrayList
		*
		* 			ArrayList -> implemented based on dynamic array.
		* 			LinkedList -> implemented based on doubly linked list.
		*
		* 			Arrays support random access but the cost of insertion and deletion is high as large number of elements
		* 			needs to be moved.
		*
		* 			While in LinkedList insertion and deletion is just changing the pointers but it does not support
		* 			random access.
		*
		 */

		//----------------------------------------------------------------------------------------------------


		/*
		*								HashMap
		*
		* 		It contains an array table of Entry<> type inside. Entry stores key, value pairs.
		* 		It contains four fields :
		* 				1. int hashCode
		* 				2. K key
		* 				3. V value
		* 				4. Entry<K, V> next;
		*
		* 		next -> it tells that Entry is a linked list, i.e. each position in an array is a bucket, and it stores
		* 				linked list.
		* 		HashMap uses zipper method to resolve conflicts.
		* 		Entries with the same hashcode gets appended to the end of the linked list.
		*
		* 		transient Entry[] table;
		*
		*
		*
		* 		How the zipper method works :
		*							HashMap<String, String> mp = new HashMap<>();
		* 							mp.put("K1", "V1");
		* 							mp.put("K2", "V2");
		* 							mp.put("K3", "V3");
		*		1. Create a hashmap with a default size of 16.
		* 		2.Insert the <K1, V1> pair.	Calculate the hashcode of K1 -> 115 and use the remainder method to obtain
		* 			the bucket. 115%16 = 3;
		* 		3. Similarly insert K2, V2 hashcode for K2 -> 118 -> 118 %16 -> 6
		* 		4. K3, V3 ---- K3 -> 118 -> 118%16 = 6
		*
		* 		The search needs to be divided into two steps.
		* 		1. Calculate the bucket where the key value pair is located.
		* 		2. Searching sequentially in the linked list where the value is located.
		*
		* 		Time complexity is equivalent to the length of the linked list.
		*
		* 		HashMap also allows to store k,v pairs where the key is null, for null hashcode() is not supported,
		* 		so HashMap uses the 0th bucket to store the key value pair whose keys are null.
		*
		* 		The new key - value pair is stored at the head of linked list rather than end of linked list.
		*
		* 		Determining the bucket subscript.
		* 				1. Many operations need to determine the subscript of the bucket where the key value pair is stored.
		* 					int hash = hash(key);
		* 					int i = indexFor(hash, table.length);
		*
		* 				Calculate the hash value.
		* 				final int hash(Object k) {
		* 					int  h = hashSeed;
		* 					if (h != 0 && k instanceof String)
		* 						return sum.misc.Hashing.stringHash32((String) k);
		*
		* 					h ^= k.hashCode();
		* 					//this function ensures that hashCodes differ only by constant multiples times at each
		* 					// bit position have a bounded number of collisions.
		* 					h ^= (h >>> 20) ^ (h >>> 12);
		* 					return h ^ (h >>> 7) ^ (h >>> 4);
		* 				}
		*
		* 				public final int hashCode() {
		* 					return Objects.hashCode(key) ^ Objects.hashCode(value);
		* 				}
		*
		* 		Modelling :
		* 			Let x = 1<<4 which is equivalent to (2^4 = 16), it has the following properties.
		* 			i)
		* 				x 	= 00010000
			* 			x-1 = 00001111
		*
		* 			ii)
		* 				y = 10110010
		* 			  x-1 = 00001111
		* 			 y & (x-1) = 00000010
		*
		* 			This is property has the same effect of doing y % x
		*
		* 		Cost of bit operation is much smaller than that of modulo operation.
		*
		* 		The last step is to take the hash value of the key and perform modulo operation with length of bucket.
		*
		* 		static int indexFor(int h, int length) {
		* 			return h & (length-1);
		* 		}
		*
		*
		* 						Capacity expansion
		*
		* 		Basics:
		* 			Suppose length of hashmap is M, number of key value pairs stored is N. If the hash function meets
		* 			uniformity requirements, then the length of each linked list is about N/M, so the search complexity
		* 			will be O(N/M)
		*
		* 			capacity : The capacity of the table is 16 by default. It must be guaranteed that cap should be
		* 						nth power of 2.
		*
		* 			size : number of key value pairs
		*
		* 			threshold : critical value of size, if this threshold is reached then expansion must be performed.
		*
		* 			load factor : threshold = (load factor * capacity)
		*
		*
		* 		Expansion : It is implemented with resize(). It should be noted that all key value pairs of old table must
		* 					be reinserted.
		*
		* 		Suppose the old capacity was 16, then new cap will be 32.
		*
		* 		NOTE : from JDK - 1.8
		* 		When the size of the linked list reaches >=8 then linked list will be converted to red black tree.
		*
		*
		*
		*
		* 							Comparison with HashTable
		*
		* 			HashTable uses synchronized for synchronization.
		* 			HashMap can store null keys.
		* 			HashMap iterators are fail fast iterators.
		* 			HashMap does not guarantee the order of elements stored.
		*
		*
		 */


		//--------------------------------------------------------------------------------------------------------

		/*
		*								ConcurrentHashMap
		*
		*		ConcurrentHashMap and HashMap are similar in implementation. Difference is ConcurrentHashMap uses
		* 		segment lock and each segment lock maintains buckets(HashEntry). Multiple threads can simultaneously
		* 		access the buckets on different segment locks. So that it has higher degree of concurrency.
		*
		*		Segment inherits from Reentrant locks.
		* 		The default concurrency level is 16, which means 16 segments are created by default.
		*
		*
		*
		* 								size operation.
		* 		Each segment maintains a count variable to count the number of key value pairs in the segment.
		*
		* 		When performing the size operation it is necessary to traverse all segments and accumulate the count.
		*		ConcurrentHashMap tries not to lock when executing the size operation. If the result of two unlock operations
		*		are consistent, then the result can be considered correct.
		*
		*		The number of attempts is defined using RETRIES_BEFORE_LOCK which is 2, and the initial value is -1.
		* 		So the number of attempts is 3.
		*
		* 		Changes in JDK 1.8
		* 			JDK 1.7 uses segmentation lock mechanism to implement concurrent update operations. The core class
		* 			is segment, which inherits from the reentrant lock ReentrantLock, and the concurrency is equal to
		* 			the number of segments.
		*
		* 			JDK 1.8 uses the CAS operation to support higher concurrency, and uses the built-in lock
		* 			synchronized when CAS operation fails.
		* 			And in this linked list will also be converted to red black tree.
		*
		*
		*
		*
		*
		* -------------------------------------LinkedHashMap--------------------------------------------------------
		*
		* 		Storage structure.
		* 			Inherited from HashMap, it has the fast look feature same as HashMap.
		*
		* 		public class LinkedHashMap<K, V> extends HashMap<K, V> implements Map<K, V>

		* 		A doubly linked list is maintained to maintain the insertion order or LRU order.
		*
		* 		transient LinkedHashMap.Entry<K, V> head;
		* 		transient LinkedHashMap.Entry<K, V> tail;
		*
		* 		accessOrder determines the order, the default is false, and the insertion order is maintained at this time.
		*
		* 		final boolean accessOrder;
		*
		*
		*
		*
		 */


	}
}