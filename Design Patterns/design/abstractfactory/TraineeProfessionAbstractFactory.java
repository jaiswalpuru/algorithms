package abstractfactory;

public class TraineeProfessionAbstractFactory extends AbstractFactory {

	@Override
	public Profession getProfession(String type) {
		if (type == null) return null;

		if (type.equals("Engineer")) return new TraineeEngineer();
		else if (type.equals("Teacher")) return new TraineeTeacher();

		return null;
	}
}
