package strategy;

// Note avoid interfaces that just force an action.
// Using interfaces just to force the creation of a method is a bad idea.

public interface Fly {
	String fly();
}

// Now we use this interface inside animal as an instance variable.
// then we are going to dynamically change it.

class CanFly implements Fly {
	@Override
	public String fly() {
		return "I can fly";
	}
}

class CantFly implements Fly{
	@Override
	public String fly() {
		return "I cannot fly";
	}
}




