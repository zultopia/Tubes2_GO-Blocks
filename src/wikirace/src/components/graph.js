import React from 'react';
import Graph from "react-graph-vis";

function replaceUnderscores(title) {
    return title.replace(/_/g, ' '); 
}

function Graf({ paths }) {
    var tempNodes = [];
    var tempEdges = [];

    paths.map((path) => {
        var len = path.length;
        for (var i = 0; i < len; i++) {
            const nodeTitle = replaceUnderscores(path[i].Title); 

            const nodeExists = tempNodes.some((node) => node.id === nodeTitle);
            if (!nodeExists) {
                tempNodes.push({
                    id: nodeTitle,
                    label: nodeTitle, 
                    title: path[i].URL, 
                });
            }

            if (i < len - 1) {
                const fromNode = replaceUnderscores(path[i].Title);
                const toNode = replaceUnderscores(path[i + 1].Title);

                const edgeExists = tempEdges.some(
                    (edge) => edge.from === fromNode && edge.to === toNode
                );
                if (!edgeExists) {
                    tempEdges.push({
                        from: fromNode,
                        to: toNode,
                        arrows: 'to', 
                        color: { color: '#848484', highlight: '#848484' },
                    });
                }
            }
        }
    });

    // Algoritma pewarnaan berdasarkan indeks
    const colors = [
        "#e6194b",
        "#3cb44b",
        "#ffe119",
        "#4363d8",
        "#f58231",
        "#911eb4",
        "#42d4f4",
        "#f032e6",
        "#a9a9a9",
        "#469990",
    ]; 
    const nodeColors = {};
    const colorByIndex = {}; 
    
    paths.forEach((path) => {
        for (let i = 0; i < path.length; i++) {
            const nodeTitle = replaceUnderscores(path[i].Title); 

            if (!colorByIndex[i]) {
                colorByIndex[i] = colors[i % colors.length]; 
            }

            nodeColors[nodeTitle] = colorByIndex[i]; 
        }
    });

    tempNodes.forEach((node) => {
        node.color = { background: nodeColors[node.id] };
    });

    const graph = {
        nodes: tempNodes,
        edges: tempEdges,
    };

    const options = {
        nodes: {
            shape: 'dot',
            size: 16,
        },
        edges: {
            arrows: {
                to: true, 
            },
            color: { color: '#848484', highlight: '#848484' }, 
        },
        physics: {
            enabled: true,
        },
        interaction: {
            navigationButtons: true,
            zoomView: true,
            hover: true,
            zoomSpeed: 0.4,
        },
    };

    return <Graph graph={graph} options={options} />;
}

export default Graf;