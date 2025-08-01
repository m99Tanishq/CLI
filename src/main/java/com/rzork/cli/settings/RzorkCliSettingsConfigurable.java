package com.rzork.cli.settings;

import com.intellij.openapi.options.Configurable;
import com.intellij.openapi.options.ConfigurationException;
import com.intellij.openapi.ui.VerticalFlowLayout;
import com.intellij.ui.components.JBLabel;
import com.intellij.ui.components.JBTextField;
import com.intellij.util.ui.FormBuilder;
import com.rzork.cli.services.RzorkCliService;
import org.jetbrains.annotations.Nls;
import org.jetbrains.annotations.Nullable;

import javax.swing.*;
import java.awt.*;

public class RzorkCliSettingsConfigurable implements Configurable {
    
    private JBTextField apiKeyField;
    private JBTextField modelField;
    private JBTextField baseUrlField;
    private JPanel mainPanel;
    
    @Nls(capitalization = Nls.Capitalization.Title)
    @Override
    public String getDisplayName() {
        return "Rzork CLI Settings";
    }
    
    @Override
    public JComponent getPreferredFocusedComponent() {
        return apiKeyField;
    }
    
    @Nullable
    @Override
    public JComponent createComponent() {
        apiKeyField = new JBTextField();
        modelField = new JBTextField();
        baseUrlField = new JBTextField();
        
        mainPanel = FormBuilder.createFormBuilder()
                .addLabeledComponent(new JBLabel("API Key: "), apiKeyField, 1, false)
                .addLabeledComponent(new JBLabel("Model: "), modelField, 1, false)
                .addLabeledComponent(new JBLabel("Base URL: "), baseUrlField, 1, false)
                .addComponentFillVertically(new JPanel(), 0)
                .getPanel();
        
        return mainPanel;
    }
    
    @Override
    public boolean isModified() {
        RzorkCliService service = RzorkCliService.getInstance();
        return !apiKeyField.getText().equals(service.getApiKey()) ||
               !modelField.getText().equals(service.getModel()) ||
               !baseUrlField.getText().equals(service.getBaseUrl());
    }
    
    @Override
    public void apply() throws ConfigurationException {
        RzorkCliService service = RzorkCliService.getInstance();
        service.setApiKey(apiKeyField.getText());
        service.setModel(modelField.getText());
        service.setBaseUrl(baseUrlField.getText());
    }
    
    @Override
    public void reset() {
        RzorkCliService service = RzorkCliService.getInstance();
        apiKeyField.setText(service.getApiKey());
        modelField.setText(service.getModel());
        baseUrlField.setText(service.getBaseUrl());
    }
    
    @Override
    public void disposeUIResources() {
        mainPanel = null;
    }
} 