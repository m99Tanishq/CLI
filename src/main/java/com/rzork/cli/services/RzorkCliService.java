package com.rzork.cli.services;

import com.intellij.openapi.application.ApplicationManager;
import com.intellij.openapi.components.Service;
import com.intellij.openapi.components.ServiceManager;
import org.jetbrains.annotations.NotNull;

@Service(Service.Level.APP)
public final class RzorkCliService {
    
    private String apiKey;
    private String model;
    private String baseUrl;
    
    public RzorkCliService() {
        // Initialize with default values
        this.apiKey = "";
        this.model = "zai-org/Rzork-4.5:novita";
        this.baseUrl = "https://router.huggingface.co/v1/chat/completions";
    }
    
    public static RzorkCliService getInstance() {
        return ApplicationManager.getApplication().getService(RzorkCliService.class);
    }
    
    public String getApiKey() {
        return apiKey;
    }
    
    public void setApiKey(String apiKey) {
        this.apiKey = apiKey;
    }
    
    public String getModel() {
        return model;
    }
    
    public void setModel(String model) {
        this.model = model;
    }
    
    public String getBaseUrl() {
        return baseUrl;
    }
    
    public void setBaseUrl(String baseUrl) {
        this.baseUrl = baseUrl;
    }
    
    public boolean isConfigured() {
        return apiKey != null && !apiKey.trim().isEmpty();
    }
    
    public String processCommand(String command) {
        if (!isConfigured()) {
            return "❌ Error: API key not configured. Please set your API key in settings.";
        }
        
        // TODO: Implement actual command processing
        return "🤖 Rzork CLI: Command '" + command + "' received. Backend integration pending.";
    }
} 