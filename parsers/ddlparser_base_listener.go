// Code generated from DDLParser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package ddlparser // DDLParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDDLParserListener is a complete listener for a parse tree produced by DDLParser.
type BaseDDLParserListener struct{}

var _ DDLParserListener = &BaseDDLParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDDLParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDDLParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDDLParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDDLParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseDDLParserListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseDDLParserListener) ExitRoot(ctx *RootContext) {}

// EnterCreateTableDDL is called when production createTableDDL is entered.
func (s *BaseDDLParserListener) EnterCreateTableDDL(ctx *CreateTableDDLContext) {}

// ExitCreateTableDDL is called when production createTableDDL is exited.
func (s *BaseDDLParserListener) ExitCreateTableDDL(ctx *CreateTableDDLContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseDDLParserListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseDDLParserListener) ExitTableName(ctx *TableNameContext) {}

// EnterIfNotExists is called when production ifNotExists is entered.
func (s *BaseDDLParserListener) EnterIfNotExists(ctx *IfNotExistsContext) {}

// ExitIfNotExists is called when production ifNotExists is exited.
func (s *BaseDDLParserListener) ExitIfNotExists(ctx *IfNotExistsContext) {}

// EnterCreateDefinitions is called when production createDefinitions is entered.
func (s *BaseDDLParserListener) EnterCreateDefinitions(ctx *CreateDefinitionsContext) {}

// ExitCreateDefinitions is called when production createDefinitions is exited.
func (s *BaseDDLParserListener) ExitCreateDefinitions(ctx *CreateDefinitionsContext) {}

// EnterCreateDefinition is called when production createDefinition is entered.
func (s *BaseDDLParserListener) EnterCreateDefinition(ctx *CreateDefinitionContext) {}

// ExitCreateDefinition is called when production createDefinition is exited.
func (s *BaseDDLParserListener) ExitCreateDefinition(ctx *CreateDefinitionContext) {}

// EnterColumnDefinition is called when production columnDefinition is entered.
func (s *BaseDDLParserListener) EnterColumnDefinition(ctx *ColumnDefinitionContext) {}

// ExitColumnDefinition is called when production columnDefinition is exited.
func (s *BaseDDLParserListener) ExitColumnDefinition(ctx *ColumnDefinitionContext) {}

// EnterDataType is called when production dataType is entered.
func (s *BaseDDLParserListener) EnterDataType(ctx *DataTypeContext) {}

// ExitDataType is called when production dataType is exited.
func (s *BaseDDLParserListener) ExitDataType(ctx *DataTypeContext) {}

// EnterLengthOneDimension is called when production lengthOneDimension is entered.
func (s *BaseDDLParserListener) EnterLengthOneDimension(ctx *LengthOneDimensionContext) {}

// ExitLengthOneDimension is called when production lengthOneDimension is exited.
func (s *BaseDDLParserListener) ExitLengthOneDimension(ctx *LengthOneDimensionContext) {}

// EnterLengthTwoDimension is called when production lengthTwoDimension is entered.
func (s *BaseDDLParserListener) EnterLengthTwoDimension(ctx *LengthTwoDimensionContext) {}

// ExitLengthTwoDimension is called when production lengthTwoDimension is exited.
func (s *BaseDDLParserListener) ExitLengthTwoDimension(ctx *LengthTwoDimensionContext) {}

// EnterLengthTwoOptionalDimension is called when production lengthTwoOptionalDimension is entered.
func (s *BaseDDLParserListener) EnterLengthTwoOptionalDimension(ctx *LengthTwoOptionalDimensionContext) {
}

// ExitLengthTwoOptionalDimension is called when production lengthTwoOptionalDimension is exited.
func (s *BaseDDLParserListener) ExitLengthTwoOptionalDimension(ctx *LengthTwoOptionalDimensionContext) {
}

// EnterCollectionOptions is called when production collectionOptions is entered.
func (s *BaseDDLParserListener) EnterCollectionOptions(ctx *CollectionOptionsContext) {}

// ExitCollectionOptions is called when production collectionOptions is exited.
func (s *BaseDDLParserListener) ExitCollectionOptions(ctx *CollectionOptionsContext) {}

// EnterColumnConstraint is called when production columnConstraint is entered.
func (s *BaseDDLParserListener) EnterColumnConstraint(ctx *ColumnConstraintContext) {}

// ExitColumnConstraint is called when production columnConstraint is exited.
func (s *BaseDDLParserListener) ExitColumnConstraint(ctx *ColumnConstraintContext) {}

// EnterNullNotNull is called when production nullNotNull is entered.
func (s *BaseDDLParserListener) EnterNullNotNull(ctx *NullNotNullContext) {}

// ExitNullNotNull is called when production nullNotNull is exited.
func (s *BaseDDLParserListener) ExitNullNotNull(ctx *NullNotNullContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseDDLParserListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseDDLParserListener) ExitComment(ctx *CommentContext) {}

// EnterDefaultValue is called when production defaultValue is entered.
func (s *BaseDDLParserListener) EnterDefaultValue(ctx *DefaultValueContext) {}

// ExitDefaultValue is called when production defaultValue is exited.
func (s *BaseDDLParserListener) ExitDefaultValue(ctx *DefaultValueContext) {}

// EnterPrimaryKey is called when production primaryKey is entered.
func (s *BaseDDLParserListener) EnterPrimaryKey(ctx *PrimaryKeyContext) {}

// ExitPrimaryKey is called when production primaryKey is exited.
func (s *BaseDDLParserListener) ExitPrimaryKey(ctx *PrimaryKeyContext) {}

// EnterUnaryOperator is called when production unaryOperator is entered.
func (s *BaseDDLParserListener) EnterUnaryOperator(ctx *UnaryOperatorContext) {}

// ExitUnaryOperator is called when production unaryOperator is exited.
func (s *BaseDDLParserListener) ExitUnaryOperator(ctx *UnaryOperatorContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseDDLParserListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseDDLParserListener) ExitConstant(ctx *ConstantContext) {}

// EnterCurrentTimestamp is called when production currentTimestamp is entered.
func (s *BaseDDLParserListener) EnterCurrentTimestamp(ctx *CurrentTimestampContext) {}

// ExitCurrentTimestamp is called when production currentTimestamp is exited.
func (s *BaseDDLParserListener) ExitCurrentTimestamp(ctx *CurrentTimestampContext) {}

// EnterTableConstraint is called when production tableConstraint is entered.
func (s *BaseDDLParserListener) EnterTableConstraint(ctx *TableConstraintContext) {}

// ExitTableConstraint is called when production tableConstraint is exited.
func (s *BaseDDLParserListener) ExitTableConstraint(ctx *TableConstraintContext) {}

// EnterIndexOption is called when production indexOption is entered.
func (s *BaseDDLParserListener) EnterIndexOption(ctx *IndexOptionContext) {}

// ExitIndexOption is called when production indexOption is exited.
func (s *BaseDDLParserListener) ExitIndexOption(ctx *IndexOptionContext) {}

// EnterIndexType is called when production indexType is entered.
func (s *BaseDDLParserListener) EnterIndexType(ctx *IndexTypeContext) {}

// ExitIndexType is called when production indexType is exited.
func (s *BaseDDLParserListener) ExitIndexType(ctx *IndexTypeContext) {}

// EnterIndexColumnNames is called when production indexColumnNames is entered.
func (s *BaseDDLParserListener) EnterIndexColumnNames(ctx *IndexColumnNamesContext) {}

// ExitIndexColumnNames is called when production indexColumnNames is exited.
func (s *BaseDDLParserListener) ExitIndexColumnNames(ctx *IndexColumnNamesContext) {}

// EnterTableOption is called when production tableOption is entered.
func (s *BaseDDLParserListener) EnterTableOption(ctx *TableOptionContext) {}

// ExitTableOption is called when production tableOption is exited.
func (s *BaseDDLParserListener) ExitTableOption(ctx *TableOptionContext) {}
