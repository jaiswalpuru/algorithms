package solid.isp;

public interface PizzaApp {
	void acceptOrderOnline();
	void acceptWalkinOrders();
	void acceptTelephonicOrders();

	//payments
	void acceptPaymentOnline();
	void acceptCash();
}
