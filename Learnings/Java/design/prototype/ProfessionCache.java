package prototype;

import java.util.HashMap;

public class ProfessionCache {
	private static HashMap<Integer, Profession> professionMap = new HashMap<>();
	public static Profession getCloneNewProfession(int id) {
		Profession cachedInstance = professionMap.get(id);
		return cachedInstance.clone();
	}

	public static void loadProfessionCache() {
		Doctor doc = new Doctor();
		doc.id = 1;
		doc.name = "Doctor";
		professionMap.put(doc.id, doc);

		Teacher t = new Teacher();
		t.id = 2;
		t.name = "Teacher";
		professionMap.put(t.id, t);

		Engineer e = new Engineer();
		e.id = 3;
		e.name = "Engineer";
		professionMap.put(e.id, e);
	}

}
