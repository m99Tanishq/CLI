package com.rzork.cli.components;

import com.intellij.openapi.components.ProjectComponent;
import com.intellij.openapi.diagnostic.Logger;
import com.intellij.openapi.project.Project;
import com.rzork.cli.services.RzorkCliProjectService;
import org.jetbrains.annotations.NotNull;

public class RzorkCliProjectComponent implements ProjectComponent {
    
    private static final Logger LOG = Logger.getInstance(RzorkCliProjectComponent.class);
    private final Project project;
    private final RzorkCliProjectService projectService;
    
    public RzorkCliProjectComponent(@NotNull Project project) {
        this.project = project;
        this.projectService = project.getService(RzorkCliProjectService.class);
    }
    
    @Override
    public void projectOpened() {
        LOG.info("📁 Rzork CLI Project Component opened for project: " + project.getName());
        LOG.info("Project path: " + projectService.getProjectPath());
    }
    
    @Override
    public void projectClosed() {
        LOG.info("📁 Rzork CLI Project Component closed for project: " + project.getName());
    }
    
    @Override
    public void initComponent() {
        LOG.info("🚀 Rzork CLI Project Component initialized");
    }
    
    @Override
    public void disposeComponent() {
        LOG.info("🔄 Rzork CLI Project Component disposed");
    }
    
    @NotNull
    @Override
    public String getComponentName() {
        return "RzorkCliProjectComponent";
    }
} 