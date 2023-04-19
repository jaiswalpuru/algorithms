package solid.isp;

// now if we see here even if the PizzaApp does not accept telephonic orders also
// we will have to implement it.
// Clients are forced to implement unnecessary methods.

// now we create two interfaces which is for offline and online separate,
// now I am a client which deals with the offline orders only so I will just implement those.

public class PizzaClient implements PizzaAppOffline {

	@Override
	public void acceptWalkinOrders() {

	}

	@Override
	public void accpetOnlyCash() {

	}

}
