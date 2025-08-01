package com.rzork.cli.components;

import com.intellij.openapi.application.ApplicationManager;
import com.intellij.openapi.components.ApplicationComponent;
import com.intellij.openapi.diagnostic.Logger;
import com.rzork.cli.services.RzorkCliService;
import org.jetbrains.annotations.NotNull;

public class RzorkCliApplicationComponent implements ApplicationComponent {
    
    private static final Logger LOG = Logger.getInstance(RzorkCliApplicationComponent.class);
    
    @Override
    public void initComponent() {
        LOG.info("🚀 Rzork CLI Application Component initialized");
        
        // Initialize the service
        RzorkCliService service = RzorkCliService.getInstance();
        LOG.info("Rzork CLI Service initialized with model: " + service.getModel());
    }
    
    @Override
    public void disposeComponent() {
        LOG.info("🔄 Rzork CLI Application Component disposed");
    }
    
    @NotNull
    @Override
    public String getComponentName() {
        return "RzorkCliApplicationComponent";
    }
} 