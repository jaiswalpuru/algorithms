package polymorphism;

public class Monkey implements Living{
	String name;
	@Override
	public void setName(String name) {
		this.name = name;
	}

	@Override
	public String getName() {
		return null;
	}

	@Override
	public void setHeight(int height) {

	}

	@Override
	public int getHeight() {
		return 0;
	}
}
