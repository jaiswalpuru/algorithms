package prototype;

public class Main {
	public static void main(String[] args) {
		ProfessionCache.loadProfessionCache();

		Profession doc = ProfessionCache.getCloneNewProfession(1);
		System.out.println(doc);

		Profession teacher = ProfessionCache.getCloneNewProfession(2);
		System.out.println(teacher);

		Profession engineer = ProfessionCache.getCloneNewProfession(3);
		System.out.println(engineer);

		Profession doc2 = ProfessionCache.getCloneNewProfession(1);
		System.out.println(doc2);

	}
}
