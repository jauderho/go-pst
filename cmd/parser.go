// This file is part of go-pst (https://github.com/mooijtech/go-pst)
// Copyright (C) 2021 Marten Mooij (https://www.mooijtech.com/)
package main

import (
	pst "github.com/mooijtech/go-pst/pkg"
	log "github.com/sirupsen/logrus"
)

func main() {
	pstFile := pst.New("data/support.pst")

	log.Infof("Parsing file: %s", pstFile.Filepath)

	isValidSignature, err := pstFile.IsValidSignature()

	if err != nil {
		log.Errorf("Failed to read signature: %s", err)
		return
	}

	if !isValidSignature {
		log.Errorf("Invalid file signature.")
		return
	}

	contentType, err := pstFile.GetContentType()

	if err != nil {
		log.Errorf("Failed to get content type: %s", err)
		return
	}

	log.Infof("Content type: %s", contentType)

	formatType, err := pstFile.GetFormatType()

	if err != nil {
		log.Errorf("Failed to get format type: %s", err)
		return
	}

	log.Infof("Format type: %s", formatType)

	encryptionType, err := pstFile.GetEncryptionType(formatType)

	if err != nil {
		log.Errorf("Failed to get encryption type: %s", err)
		return
	}

	log.Infof("Encryption type: %s", encryptionType)

	nodeBTreeOffset, err := pstFile.GetNodeBTreeOffset(formatType)

	if err != nil {
		log.Errorf("Failed to get node b-tree offset: %s", err)
		return
	}

	log.Infof("Node b-tree offset: %d", nodeBTreeOffset)

	blockBTreeOffset, err := pstFile.GetBlockBTreeOffset(formatType)

	log.Infof("Block b-tree offset: %d", blockBTreeOffset)

	btreeNodeEntryCount, err := pstFile.GetBTreeNodeEntryCount(nodeBTreeOffset, formatType)

	if err != nil {
		log.Errorf("Failed to get b-tree entry count: %s", err)
		return
	}

	log.Infof("Node b-tree entry count: %d", btreeNodeEntryCount)

	btreeNodeEntrySize, err := pstFile.GetBTreeNodeEntrySize(nodeBTreeOffset, formatType)

	if err != nil {
		log.Infof("Failed to get node b-tree entry size: %s", err)
		return
	}

	log.Infof("Node b-tree entry size: %d", btreeNodeEntrySize)

	btreeNodeLevel, err := pstFile.GetBTreeNodeLevel(nodeBTreeOffset, formatType)

	if err != nil {
		log.Errorf("Failed to get node b-tree level: %s", btreeNodeLevel)
	}

	log.Infof("Node b-tree level: %d", btreeNodeLevel)

	btreeNodeEntries, err := pstFile.GetBTreeNodeEntries(nodeBTreeOffset, formatType)

	if err != nil {
		log.Errorf("Failed to get node b-tree entries: %s", err)
		return
	}

	log.Infof("Node b-tree entries: %d", len(btreeNodeEntries))

	rootFolderNode, err := pstFile.FindBTreeNode(nodeBTreeOffset, 290, formatType)

	if err != nil {
		log.Infof("Failed to find root folder node: %s", err)
		return
	}

	log.Infof("Root folder node: %b", rootFolderNode.Data)

	rootFolderNodeDataIdentifier, err := rootFolderNode.GetDataIdentifier(formatType)

	if err != nil {
		log.Errorf("Failed to get root folder node data identifier: %s", err)
		return
	}

	log.Infof("Root folder node data identifier: %d", rootFolderNodeDataIdentifier)

	rootFolderNodeDataNode, err := pstFile.FindBTreeNode(blockBTreeOffset, rootFolderNodeDataIdentifier, formatType)

	if err != nil {
		log.Errorf("Failed to get root folder node data node: %s", err)
		return
	}

	log.Infof("Root folder node data node: %b", rootFolderNodeDataNode.Data)

	rootFolderNodeDataNodeOffset, err := rootFolderNodeDataNode.GetFileOffset(false, formatType)

	if err != nil {
		log.Errorf("Failed to get root folder node data node offset: %s", err)
		return
	}

	log.Infof("Root folder node data node offset: %d", rootFolderNodeDataNodeOffset)

	rootFolderNodeDataNodeSize, err := rootFolderNodeDataNode.GetSize(formatType)

	if err != nil {
		log.Errorf("Failed to get root folder node data node size: %s", err)
		return
	}

	log.Infof("Root folder node data node size: %d", rootFolderNodeDataNodeSize)

	rootFolderNodeDataNodeHeapOnNode, err := pstFile.GetHeapOnNode(rootFolderNodeDataNode, formatType)

	if err != nil {
		log.Errorf("Failed to decrypt table: %s", err)
		return
	}

	log.Infof("Has valid block signature: %t", rootFolderNodeDataNodeHeapOnNode.IsValidSignature())
	log.Infof("Table type: %d", rootFolderNodeDataNodeHeapOnNode.GetTableType())

	err = pstFile.GetRootFolder(formatType)

	if err != nil {
		log.Errorf("Failed to get root folder: %s", err)
		return
	}
}
