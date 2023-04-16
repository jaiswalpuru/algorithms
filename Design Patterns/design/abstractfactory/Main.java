package abstractfactory;

public class Main {
	public static void main(String[] args) {
		AbstractFactory af = AbstractFactoryProducer.getProfession(true);
		Profession p = af.getProfession("Engineer");
		p.print();
		AbstractFactory af2 = AbstractFactoryProducer.getProfession(false);
		Profession p1 = af2.getProfession("Teacher");
		p1.print();
	}
}
