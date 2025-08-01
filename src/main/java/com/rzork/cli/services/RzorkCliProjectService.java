package com.rzork.cli.services;

import com.intellij.openapi.components.Service;
import com.intellij.openapi.project.Project;
import org.jetbrains.annotations.NotNull;

@Service(Service.Level.PROJECT)
public final class RzorkCliProjectService {
    
    private final Project project;
    private String projectPath;
    private boolean isIndexed;
    
    public RzorkCliProjectService(@NotNull Project project) {
        this.project = project;
        this.projectPath = project.getBasePath();
        this.isIndexed = false;
    }
    
    public Project getProject() {
        return project;
    }
    
    public String getProjectPath() {
        return projectPath;
    }
    
    public void setProjectPath(String projectPath) {
        this.projectPath = projectPath;
    }
    
    public boolean isIndexed() {
        return isIndexed;
    }
    
    public void setIndexed(boolean indexed) {
        isIndexed = indexed;
    }
    
    public String getProjectName() {
        return project.getName();
    }
    
    public void indexProject() {
        // TODO: Implement project indexing
        this.isIndexed = true;
    }
    
    public void clearIndex() {
        // TODO: Implement index clearing
        this.isIndexed = false;
    }
} 