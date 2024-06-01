describe('template spec', () => {
  it('passes', () => {
    cy.visit('/')
  })

  it('renders correctly', () => {
    cy.visit('/management')
    cy.get('h1').should('contain', 'PagSeguro')
    cy.get('[data-testid="kanban-column"]').should('have.length', 5);

  })
  
  it('changes view mode', () => {
    cy.visit('/management')
    cy.get('[data-testid="history-view"]').click()
    cy.get('[data-testid="history-base"]').should('exist')
  })

  it('drags and drops cards', () => {
    cy.visit('/management')
    
    cy.get('[data-testid="kanban-column"]').eq(0).as('sourceColumn')
    cy.get('[data-testid="kanban-column"]').eq(1).as('destinationColumn')
  
    cy.get('@sourceColumn').find('[data-testid="kanban-card"]').eq(0).as('draggedCard')
    cy.get('@destinationColumn').as('destinationDroppable')

    cy.get('@draggedCard').trigger('dragstart')

    cy.get('@destinationColumn').trigger('drop')

    cy.get('@draggedCard').trigger('dragend')
  
    cy.get('@destinationDroppable').find('[data-testid="kanban-card"]').should('have.length', 1)
  })

  it('selects view mode', () => {
    cy.visit('/management')
    cy.get('[data-testid="history-view"]').click()
    cy.get('[data-testid="management-view"]').click()
    cy.get('[data-testid="kanban-base"]').should('exist')
  })
})

