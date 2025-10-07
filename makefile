PATH_MIGRATION=./internal/infrastructure/database/migrations
DOCKER ?= docker
DOCKER_COMPOSE = $(DOCKER) compose -f compose.yml


dev:
	@echo "Running..."
	go run cmd/main.go

stack-up:
	$(DOCKER_COMPOSE) up -d


	
.PHONY: migration migrate-up migrate-down migrate-version migrate-force

migration-apply:
	go run cmd/migrate/main.go

# Créer une nouvelle migration
# Usage: make migration NAME=create_users_table
migration:
	@if [ "$(NAME)" = "" ]; then \
		echo "Erreur: Spécifiez un nom avec NAME=..."; \
		echo "Usage: make migration NAME=create_users_table"; \
		exit 1; \
	fi
	@echo "Création de la migration: $(NAME)"
	migrate create -ext sql -dir $(PATH_MIGRATION) -seq $(NAME)
	@echo "Migration créée dans $(PATH_MIGRATION)"

# Appliquer toutes les migrations
migrate-up:
	@echo "Application des migrations..."
	migrate -database "$(DB_URL)" -path $(PATH_MIGRATION) up
	@echo "Migrations appliquées"

# Revenir en arrière d'une migration
migrate-down:
	@echo "Rollback d'une migration..."
	migrate -database "$(DB_URL)" -path $(PATH_MIGRATION) down 1
	@echo "Rollback effectué"

# Voir la version actuelle
migrate-version:
	@echo "Version actuelle des migrations:"
	migrate -database "$(DB_URL)" -path $(PATH_MIGRATION) version

# Forcer une version (en cas de problème)
# Usage: make migrate-force VERSION=1
migrate-force:
	@if [ "$(VERSION)" = "" ]; then \
		echo "Erreur: Spécifiez une version avec VERSION=..."; \
		echo "Usage: make migrate-force VERSION=1"; \
		exit 1; \
	fi
	@echo "🔧 Force de la version $(VERSION)..."
	migrate -database "$(DB_URL)" -path $(PATH_MIGRATION) force $(VERSION)
	@echo "Version forcée à $(VERSION)"

# Créer le dossier migrations s'il n'existe pas
init-migrations:
	@echo "Création du dossier $(PATH_MIGRATION)..."
	@mkdir -p $(PATH_MIGRATION)
	@echo "Dossier créé"

reset-database:
	migrate -database "$(DB_URL)" -path "$(PATH_MIGRATION)" drop

# Aide
help:
	@echo "Commandes disponibles:"
	@echo ""
	@echo "Gestion des migrations:"
	@echo "  make migration NAME=create_users_table  - Créer une nouvelle migration"
	@echo "  make migrate-up                         - Appliquer toutes les migrations"
	@echo "  make migrate-down                       - Revenir en arrière d'une migration"
	@echo "  make migrate-version                    - Voir la version actuelle"
	@echo "  make migrate-force VERSION=1            - Forcer une version"
	@echo "  make init-migrations                    - Créer le dossier migrations"
	@echo ""
	@echo "Configuration:"
	@echo "  DB_URL: $(DB_URL)"
	@echo "  PATH_MIGRATION: $(PATH_MIGRATION)"
