package abstractfactory;

public class ProfessionAbstractFactory extends AbstractFactory {

	@Override
	public Profession getProfession(String type) {
		if (type == null) return null;

		if (type.equals("Teacher")) {
			return  new Teacher();
		} else if (type.equals("Engineer")) {
			return new Engineer();
		}

		return null;
	}
}
