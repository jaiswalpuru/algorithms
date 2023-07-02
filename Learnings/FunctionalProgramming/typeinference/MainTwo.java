package typeinference;

import javax.swing.*;
import java.awt.*;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

public class MainTwo {
	private final static int FIELD_WIDTH = 20;
	private static JTextField staticTextField;
	public static void main(String[] args) {
		JFrame frame = new JFrame();

		staticTextField = new JTextField(FIELD_WIDTH);
		staticTextField.setText("Static Field");

		JTextField localTextField = new JTextField(FIELD_WIDTH);
		localTextField.setText("Local variable");

		JButton helloButton = new JButton("Say Hello");
		helloButton.addActionListener(new ActionListener() {
			@Override
			public void actionPerformed(ActionEvent actionEvent) {
				staticTextField.setText("Hello, World!");
				localTextField.setText("Hello, World!");
			}
		});

		JButton goodByButton = new JButton("GoodBye");
		goodByButton.addActionListener(actionEvent -> {
			staticTextField.setText("Goodbye, World!");
			localTextField.setText("Goodbye, World!");
		});

		Container contentPane = frame.getContentPane();
		contentPane.setLayout(new FlowLayout());
		contentPane.add(staticTextField);
		contentPane.add(localTextField);
		contentPane.add(helloButton);
		contentPane.add(goodByButton);

//		staticTextField = null;
//		localTextField = null;

		frame.setAlwaysOnTop(true);
		frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
		frame.pack();
		frame.setVisible(true);
	}
}
