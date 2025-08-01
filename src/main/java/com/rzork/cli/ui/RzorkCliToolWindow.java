package com.rzork.cli.ui;

import com.intellij.openapi.project.Project;
import com.intellij.openapi.wm.ToolWindow;
import com.intellij.ui.components.JBScrollPane;
import com.intellij.ui.components.JBTextArea;
import com.intellij.ui.components.JBTextField;
import com.intellij.util.ui.JBUI;

import javax.swing.*;
import java.awt.*;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

public class RzorkCliToolWindow {
    private final Project project;
    private final ToolWindow toolWindow;
    private final JPanel mainPanel;
    private final JBTextArea outputArea;
    private final JBTextField inputField;
    private final JButton sendButton;
    private final JButton clearButton;
    private final JButton settingsButton;

    public RzorkCliToolWindow(Project project, ToolWindow toolWindow) {
        this.project = project;
        this.toolWindow = toolWindow;
        
        // Create main panel
        mainPanel = new JPanel(new BorderLayout());
        mainPanel.setBorder(JBUI.Borders.empty(10));
        
        // Create output area
        outputArea = new JBTextArea();
        outputArea.setEditable(false);
        outputArea.setFont(new Font("Monospaced", Font.PLAIN, 12));
        outputArea.setBackground(new Color(43, 43, 43));
        outputArea.setForeground(new Color(187, 187, 187));
        outputArea.setCaretColor(new Color(187, 187, 187));
        
        JBScrollPane scrollPane = new JBScrollPane(outputArea);
        scrollPane.setPreferredSize(new Dimension(600, 400));
        
        // Create input panel
        JPanel inputPanel = new JPanel(new BorderLayout());
        inputPanel.setBorder(JBUI.Borders.empty(5, 0));
        
        inputField = new JBTextField();
        inputField.setFont(new Font("Monospaced", Font.PLAIN, 12));
        inputField.addActionListener(new ActionListener() {
            @Override
            public void actionPerformed(ActionEvent e) {
                sendMessage();
            }
        });
        
        // Create button panel
        JPanel buttonPanel = new JPanel(new FlowLayout(FlowLayout.LEFT));
        
        sendButton = new JButton("Send");
        sendButton.addActionListener(new ActionListener() {
            @Override
            public void actionPerformed(ActionEvent e) {
                sendMessage();
            }
        });
        
        clearButton = new JButton("Clear");
        clearButton.addActionListener(new ActionListener() {
            @Override
            public void actionPerformed(ActionEvent e) {
                clearOutput();
            }
        });
        
        settingsButton = new JButton("Settings");
        settingsButton.addActionListener(new ActionListener() {
            @Override
            public void actionPerformed(ActionEvent e) {
                openSettings();
            }
        });
        
        buttonPanel.add(sendButton);
        buttonPanel.add(clearButton);
        buttonPanel.add(settingsButton);
        
        inputPanel.add(inputField, BorderLayout.CENTER);
        inputPanel.add(buttonPanel, BorderLayout.EAST);
        
        // Add components to main panel
        mainPanel.add(scrollPane, BorderLayout.CENTER);
        mainPanel.add(inputPanel, BorderLayout.SOUTH);
        
        // Add welcome message
        appendOutput("🚀 Welcome to Rzork CLI Assistant!\n");
        appendOutput("Type your questions or commands below.\n");
        appendOutput("Use 'help' to see available commands.\n\n");
    }
    
    public JComponent getContent() {
        return mainPanel;
    }
    
    private void sendMessage() {
        String message = inputField.getText().trim();
        if (message.isEmpty()) {
            return;
        }
        
        appendOutput("You: " + message + "\n");
        inputField.setText("");
        
        // TODO: Integrate with Rzork CLI backend
        appendOutput("🤖 Rzork CLI: Processing your request...\n");
        appendOutput("(This is a placeholder response. Backend integration pending)\n\n");
    }
    
    private void clearOutput() {
        outputArea.setText("");
        appendOutput("🚀 Welcome to Rzork CLI Assistant!\n");
        appendOutput("Type your questions or commands below.\n");
        appendOutput("Use 'help' to see available commands.\n\n");
    }
    
    private void openSettings() {
        // TODO: Open settings dialog
        appendOutput("⚙️ Settings dialog will be implemented soon.\n\n");
    }
    
    private void appendOutput(String text) {
        SwingUtilities.invokeLater(() -> {
            outputArea.append(text);
            outputArea.setCaretPosition(outputArea.getDocument().getLength());
        });
    }
} 