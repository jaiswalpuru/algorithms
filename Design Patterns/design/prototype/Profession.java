package prototype;

public abstract class Profession implements Cloneable{

	public int id;
	public String name;
	abstract void print();

	@Override
	public Profession clone() {
		try {
			return (Profession) super.clone();
		} catch (CloneNotSupportedException e) {
			e.printStackTrace();
			throw new AssertionError();
		}
	}
}
