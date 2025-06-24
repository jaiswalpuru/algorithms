package factory;

public class ProfessionFactory {
	public Profession getProfession(String type) {
		if (type == null) return null;

		switch (type) {
			case "Doctor":
				return new Doctor();
			case "Engineer":
				return new Engineer();
			case "Teacher":
				return new Teacher();
		}

		return null;
	}
}
